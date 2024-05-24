package response

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	timeFormat = "2006-01-02 15:04:05"
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

func (r *MySqlRepo) SaveResponse(ctx context.Context, resp Response) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction failed: %w", err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO responses (id, form_id, form_version, user_id, submitted_at) VALUES (?, ?, ?, ?, ?)",
		resp.Id, resp.FormId, resp.FormVersion, nil, resp.SubmittedAt)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Printf("rollback transaction failed: %v", err)
		}
		return fmt.Errorf("insert form failed: %w", err)
	}

	for _, a := range resp.Answers {
		if err := r.saveAnswer(ctx, tx, resp.Id, a); err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("rollback transaction failed: %v", err)
			}
			return fmt.Errorf("save answer failed: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction failed: %w", err)
	}

	return nil
}

func (r *MySqlRepo) saveAnswer(ctx context.Context, tx *sql.Tx, responseId string, answer Answer) error {
	switch a := answer.(type) {
	case TextAnswer:
		_, err := tx.ExecContext(ctx, "INSERT INTO answers (response_id, question_order, answer_text) VALUES (?, ?, ?)",
			responseId, a.QuestionOrder, a.Value)
		if err != nil {
			return fmt.Errorf("insert text answer failed: %w", err)
		}

	case RadioAnswer:
		_, err := tx.ExecContext(ctx, "INSERT INTO answers (response_id, question_order, answer_text) VALUES (?, ?, ?)",
			responseId, a.QuestionOrder, a.Value)
		if err != nil {
			return fmt.Errorf("insert radio answer failed: %w", err)
		}

	case CheckboxAnswer:
		for _, v := range a.Values {
			_, err := tx.ExecContext(ctx, "INSERT INTO answers (response_id, question_order, answer_text) VALUES (?, ?, ?)",
				responseId, a.QuestionOrder, v)
			if err != nil {
				return fmt.Errorf("insert checkbox answer failed: %w", err)
			}
		}
	}

	return nil
}
