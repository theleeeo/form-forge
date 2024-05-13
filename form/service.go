package form

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/models"
	"github.com/theleeeo/form-forge/repo"
)

var (
	// These variables are used to make the code testable
	UUIDNew = uuid.New
	TimeNow = time.Now
)

func NewService(repo *repo.MySqlRepo) *Service {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo *repo.MySqlRepo
}

func (s *Service) CreateNewForm(ctx context.Context, params CreateFormParams) (Form, error) {
	if params.Title == "" {
		return Form{}, fmt.Errorf("title is required")
	}

	if len(params.Questions) == 0 {
		return Form{}, fmt.Errorf("questions are required")
	}

	id := UUIDNew().String()
	form := &models.Form{
		ID:        id,
		BaseID:    id,
		Version:   1,
		Title:     params.Title,
		CreatedAt: TimeNow().UTC(),
	}

	questions := make([]models.Question, 0, len(params.Questions))
	for _, q := range params.Questions {
		var question models.Question

		base := models.QuestionBase{
			FormID: form.ID,
			Title:  q.Title,
		}

		switch q.Type {
		case models.QuestionTypeText:
			question = models.TextQuestion{
				QuestionBase: base,
			}
		case models.QuestionTypeRadio:
			if len(q.Options) == 0 {
				return Form{}, fmt.Errorf("options are required for radio questions")
			}

			question = models.RadioQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		case models.QuestionTypeCheckbox:
			if len(q.Options) == 0 {
				return Form{}, fmt.Errorf("options are required for checkbox questions")
			}

			question = models.CheckboxQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		default:
			return Form{}, fmt.Errorf("invalid question type: %d", q.Type)
		}

		questions = append(questions, question)

		if err := question.Validate(); err != nil {
			return Form{}, fmt.Errorf("question validation failed: %w", err)
		}
	}

	err := s.repo.CreateForm(ctx, form, questions)
	if err != nil {
		return Form{}, err
	}

	return Form{
		Form:      *form,
		repo:      s.repo,
		questions: questions,
	}, nil
}

func (s *Service) GetForm(ctx context.Context, id string) (Form, error) {
	form, err := s.repo.GetForm(ctx, id)
	if err != nil {
		return Form{}, err
	}

	return Form{
		Form: form,
		repo: s.repo,
	}, nil
}
