package response

import (
	"context"
	"fmt"
	"strconv"

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

func (r *Repo) SaveResponse(ctx context.Context, resp Response) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "INSERT INTO responses (id, form_version_id, submitted_at) VALUES ($1, $2, $3)",
		resp.Id, resp.FormVersionId, resp.SubmittedAt)
	if err != nil {
		return fmt.Errorf("inserting response: %w", err)
	}

	for _, a := range resp.Answers {
		if err := r.saveAnswer(ctx, tx, resp.Id, a); err != nil {
			return fmt.Errorf("saving answer: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *Repo) saveAnswer(ctx context.Context, tx pgx.Tx, responseId uuid.UUID, answer Answer) error {
	switch a := answer.(type) {
	case TextAnswer:
		_, err := tx.Exec(ctx, "INSERT INTO answers (response_id, question_id, answer_text) VALUES ($1, $2, $3)",
			responseId, a.QuestionId, a.Value)
		if err != nil {
			return fmt.Errorf("inserting text answer: %w", err)
		}

	case RadioAnswer:
		_, err := tx.Exec(ctx, "INSERT INTO answers (response_id, question_id, answer_text) VALUES ($1, $2, $3)",
			responseId, a.QuestionId, strconv.Itoa(a.Value))
		if err != nil {
			return fmt.Errorf("inserting radio answer: %w", err)
		}

	case CheckboxAnswer:
		for _, v := range a.Values {
			_, err := tx.Exec(ctx, "INSERT INTO answers (response_id, question_id, answer_text) VALUES ($1, $2, $3)",
				responseId, a.QuestionId, strconv.Itoa(v))
			if err != nil {
				return fmt.Errorf("inserting checkbox answer: %w", err)
			}
		}
	}

	return nil
}
