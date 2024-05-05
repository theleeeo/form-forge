package form

import (
	"context"
	"fmt"

	"go.leeeo.se/form-forge/models"
	"go.leeeo.se/form-forge/repo"
)

func NewForm(f models.Form) *Form {
	return &Form{
		Form: f,
	}
}

// A representation of a complete form.
type Form struct {
	repo *repo.MySqlRepo

	models.Form
	questions []models.Question
}

func (f *Form) Questions(ctx context.Context) ([]models.Question, error) {
	if f.questions != nil {
		return f.questions, nil
	}

	questions, err := f.repo.GetQuestions(ctx, f.ID)
	if err != nil {
		return nil, fmt.Errorf("get questions failed: %w", err)
	}

	f.questions = questions
	return f.questions, nil
}
