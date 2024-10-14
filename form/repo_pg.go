package form

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewPgRepo(dbpool *pgxpool.Pool) *Repo {
	return &Repo{
		conn: dbpool,
	}
}

func (r *Repo) CreateForm(ctx context.Context, form Form, questions []Question) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO forms (base_id, version_id, version, title, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		form.BaseId, form.VersionId, form.Version, form.Title, form.Description, form.CreatedAt)
	if err != nil {
		return fmt.Errorf("inserting form: %w", err)
	}

	if err := r.insertQuestions(ctx, tx, form.VersionId, questions); err != nil {
		return fmt.Errorf("inserting questions: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Repo) insertQuestions(ctx context.Context, tx pgx.Tx, formVersionId uuid.UUID, questions []Question) error {
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

		_, err := tx.Exec(ctx, "INSERT INTO questions (id, form_version_id, order_idx, title, question_type) VALUES ($1, $2, $3, $4, $5)",
			questionBase.Id, formVersionId, i, questionBase.Title, questionType)
		if err != nil {
			return err
		}

		switch q := q.(type) {
		case RadioQuestion:
			if err := r.insertOptions(ctx, tx, q.Id, q.Options); err != nil {
				return fmt.Errorf("inserting radio options: %w", err)
			}
		case CheckboxQuestion:
			if err := r.insertOptions(ctx, tx, q.Id, q.Options); err != nil {
				return fmt.Errorf("inserting checkbox options: %w", err)
			}
		}
	}

	return nil
}

func (r *Repo) insertOptions(ctx context.Context, tx pgx.Tx, questionID uuid.UUID, options []string) error {
	for i, option := range options {
		_, err := tx.Exec(ctx, "INSERT INTO options (question_id, order_idx, option_text) VALUES ($1, $2, $3)", questionID, i, option)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) GetLatestVersionOfBase(ctx context.Context, baseId uuid.UUID) (Form, error) {
	var versionID string
	err := r.conn.QueryRow(ctx, "SELECT version_id FROM forms WHERE base_id = $1 ORDER BY version DESC LIMIT 1", baseId).Scan(&versionID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return Form{}, ErrNotFound
		}
		return Form{}, err
	}

	return r.GetVersion(ctx, versionID)
}

func (r *Repo) ListForms(ctx context.Context, params ListFormsParams) ([]Form, error) {
	rows, err := r.conn.Query(ctx, `SELECT f.version_id
	FROM forms f
	INNER JOIN (
		SELECT base_id, MAX(version) AS max_version
		FROM forms
		GROUP BY base_id
	) latest_versions ON f.base_id = latest_versions.base_id AND f.version = latest_versions.max_version
	ORDER BY f.created_at DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forms []Form
	for rows.Next() {
		var version_id string
		if err := rows.Scan(&version_id); err != nil {
			return nil, err
		}

		form, err := r.GetVersion(ctx, version_id)
		if err != nil {
			return nil, err
		}

		forms = append(forms, form)
	}

	return forms, nil
}

func (r *Repo) GetVersion(ctx context.Context, versionId string) (Form, error) {
	var form Form

	err := r.conn.QueryRow(ctx, "SELECT base_id, version_id, version, title, description, created_at FROM forms WHERE version_id = $1", versionId).
		Scan(&form.BaseId, &form.VersionId, &form.Version, &form.Title, &form.Description, &form.CreatedAt)
	if err != nil {
		return Form{}, fmt.Errorf("querying form: %w", err)
	}

	return form, nil
}

func (r *Repo) GetQuestions(ctx context.Context, baseId uuid.UUID) ([]Question, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, title, question_type FROM questions WHERE form_version_id = (SELECT version_id FROM forms WHERE base_id = $1 ORDER BY version DESC LIMIT 1) ORDER BY order_idx", baseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var base QuestionBase
		var questionType QuestionType
		if err := rows.Scan(&base.Id, &base.Title, &questionType); err != nil {
			return nil, err
		}

		var q Question
		switch questionType {
		case QuestionTypeText:
			q = TextQuestion{QuestionBase: base}
		case QuestionTypeRadio:
			options, err := r.getOptions(ctx, base.Id)
			if err != nil {
				return nil, err
			}
			q = RadioQuestion{QuestionBase: base, Options: options}
		case QuestionTypeCheckbox:
			options, err := r.getOptions(ctx, base.Id)
			if err != nil {
				return nil, err
			}
			q = CheckboxQuestion{QuestionBase: base, Options: options}
		}

		questions = append(questions, q)
	}

	return questions, nil
}

func (r *Repo) getOptions(ctx context.Context, questionID uuid.UUID) ([]string, error) {
	rows, err := r.conn.Query(ctx, "SELECT option_text FROM options WHERE question_id = $1 ORDER BY order_idx", questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []string
	for rows.Next() {
		var option string
		if err := rows.Scan(&option); err != nil {
			return nil, err
		}
		options = append(options, option)
	}

	return options, nil
}

func (r *Repo) GetQuestionsOfVersion(ctx context.Context, varsionId uuid.UUID) ([]Question, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, title, question_type FROM questions WHERE form_version_id = $1 ORDER BY order_idx", varsionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var base QuestionBase
		var questionType QuestionType
		if err := rows.Scan(&base.Id, &base.Title, &questionType); err != nil {
			return nil, err
		}

		var q Question
		switch questionType {
		case QuestionTypeText:
			q = TextQuestion{QuestionBase: base}
		case QuestionTypeRadio:
			options, err := r.getOptions(ctx, base.Id)
			if err != nil {
				return nil, err
			}
			q = RadioQuestion{QuestionBase: base, Options: options}
		case QuestionTypeCheckbox:
			options, err := r.getOptions(ctx, base.Id)
			if err != nil {
				return nil, err
			}
			q = CheckboxQuestion{QuestionBase: base, Options: options}
		}

		questions = append(questions, q)
	}

	return questions, nil
}
