package form

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var (
	// ErrNotFound is returned when the requested resource is not found.
	ErrNotFound = errors.New("not found")
)

type MySqlConfig struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type MySqlRepo struct {
	db *sql.DB
}

// NewMySql creates a new MySQL repository.
// If db is nil, a new connection to the MySQL database is created using the provided config.
// If db is provided, it is used as the connection to the MySQL database.
// The caller is responsible for closing the repository using the Close method.
func NewMySql(cfg *MySqlConfig, db *sql.DB) (*MySqlRepo, error) {
	if db == nil {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.User, cfg.Password, cfg.Address, cfg.Database))
		if err != nil {
			return nil, fmt.Errorf("open mysql failed: %w", err)
		}
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping mysql failed: %w", err)
	}

	fmt.Println("Connected to MySQL")

	return &MySqlRepo{
		db: db,
	}, nil
}

func (r *MySqlRepo) Close() error {
	return r.db.Close()
}

func (r *MySqlRepo) Ping() error {
	return r.db.Ping()
}

func (r *MySqlRepo) CreateForm(ctx context.Context, form Form) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO forms (id, version, title, created_at, created_by) VALUES (?, ?, ?, ?, ?)",
		form.ID, form.Version, form.Title, form.CreatedAt, form.CreatedBy)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("rollback transaction failed: %v", err)
		}
		return fmt.Errorf("insert form failed: %w", err)
	}

	if err := r.insertQuestions(ctx, tx, form.Questions); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("rollback transaction failed: %v", err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction failed: %w", err)
	}

	return nil
}

func (r *MySqlRepo) insertQuestions(ctx context.Context, tx *sql.Tx, questions []Question) error {
	for i, q := range questions {
		questionBase := q.Question()
		var questionType QuestionType

		switch q.(type) {
		case TextQuestion:
			questionType = QuestionTypeText
		case RadioQuestion:
			questionType = QuestionTypeRadio
		case CheckboxQuestion:
			questionType = QuestionTypeCheckbox
		}

		result, err := tx.ExecContext(ctx, "INSERT INTO questions (form_id, form_version, order_idx, title, question_type) VALUES (?, ?, ?, ?, ?)",
			questionBase.FormID, questionBase.FormVersion, i, questionBase.Title, questionType)
		if err != nil {
			return fmt.Errorf("insert question failed: %w", err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("get last insert id failed: %w", err)
		}

		switch q := q.(type) {
		case RadioQuestion:
			if err := r.insertOptions(ctx, tx, id, q.Options); err != nil {
				return fmt.Errorf("insert radio options failed: %w", err)
			}
		case CheckboxQuestion:
			if err := r.insertOptions(ctx, tx, id, q.Options); err != nil {
				return fmt.Errorf("insert checkbox options failed: %w", err)
			}
		}
	}

	return nil
}

func (r *MySqlRepo) insertOptions(ctx context.Context, tx *sql.Tx, questionID int64, options []string) error {
	for i, option := range options {
		_, err := tx.ExecContext(ctx, "INSERT INTO options (question_id, order_idx, option_text) VALUES (?, ?, ?)", questionID, i, option)
		if err != nil {
			return fmt.Errorf("insert radio option failed: %w", err)
		}
	}
	return nil
}

func (r *MySqlRepo) GetForm(ctx context.Context, id string) (Form, error) {
	formBase, err := r.getFormBase(ctx, id)
	if err != nil {
		return Form{}, err
	}

	questions, err := r.GetQuestions(ctx, id)
	if err != nil {
		return Form{}, fmt.Errorf("get questions failed: %w", err)
	}

	return Form{
		FormBase:  formBase,
		Questions: questions,
	}, nil
}

func (r *MySqlRepo) getFormBase(ctx context.Context, id string) (FormBase, error) {
	var form FormBase
	var createdAt string
	err := r.db.QueryRowContext(ctx, "SELECT id, version, title, created_at, created_by FROM forms WHERE id = ?", id).
		Scan(&form.ID, &form.Version, &form.Title, &createdAt, &form.CreatedBy)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return FormBase{}, ErrNotFound
		}
		return FormBase{}, fmt.Errorf("get form failed: %w", err)
	}

	createdAtTime, err := time.Parse(timeFormat, createdAt)
	if err != nil {
		return FormBase{}, fmt.Errorf("parse created_at failed: %w", err)
	}
	form.CreatedAt = createdAtTime

	return form, nil
}

func (r *MySqlRepo) GetQuestions(ctx context.Context, formID string) ([]Question, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, form_id, form_version, title, question_type FROM questions WHERE form_id = ? ORDER BY order_idx", formID)
	if err != nil {
		return nil, fmt.Errorf("get questions failed: %w", err)
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var base QuestionBase
		var questionType QuestionType
		var id int64
		if err := rows.Scan(&id, &base.FormID, &base.FormVersion, &base.Title, &questionType); err != nil {
			return nil, fmt.Errorf("scan question failed: %w", err)
		}

		var q Question
		switch questionType {
		case QuestionTypeText:
			q = TextQuestion{QuestionBase: base}
		case QuestionTypeRadio:
			options, err := r.getOptions(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("get options failed: %w", err)
			}
			q = RadioQuestion{QuestionBase: base, Options: options}
		case QuestionTypeCheckbox:
			options, err := r.getOptions(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("get options failed: %w", err)
			}
			q = CheckboxQuestion{QuestionBase: base, Options: options}
		}

		questions = append(questions, q)
	}

	return questions, nil
}

func (r *MySqlRepo) getOptions(ctx context.Context, questionID int64) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT option_text FROM options WHERE question_id = ? ORDER BY order_idx", questionID)
	if err != nil {
		return nil, fmt.Errorf("get options failed: %w", err)
	}
	defer rows.Close()

	var options []string
	for rows.Next() {
		var option string
		if err := rows.Scan(&option); err != nil {
			return nil, fmt.Errorf("scan option failed: %w", err)
		}
		options = append(options, option)
	}

	return options, nil
}
