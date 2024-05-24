package form

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	// These variables are used to make the code testable
	UUIDNew = uuid.New
	TimeNow = time.Now
)

func NewService(repo *MySqlRepo) *Service {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo *MySqlRepo
}

type CreateFormParams struct {
	Title     string
	Questions []CreateQuestionParams
}

type CreateQuestionParams struct {
	Type  QuestionType
	Title string
	// Options is only required for radio and checkbox questions
	Options []string
}

func (s *Service) CreateNewForm(ctx context.Context, params CreateFormParams) (Form, error) {
	if params.Title == "" {
		return Form{}, fmt.Errorf("title is required")
	}

	if len(params.Questions) == 0 {
		return Form{}, fmt.Errorf("questions are required")
	}

	id := UUIDNew().String()
	form := Form{
		FormBase: FormBase{
			ID:        id,
			Version:   1,
			Title:     params.Title,
			CreatedAt: TimeNow().UTC(),
		},
	}

	questions := make([]Question, 0, len(params.Questions))
	for _, q := range params.Questions {
		var question Question

		base := QuestionBase{
			FormID:      form.ID,
			FormVersion: form.Version,
			Title:       q.Title,
		}

		switch q.Type {
		case QuestionTypeText:
			question = TextQuestion{
				QuestionBase: base,
			}
		case QuestionTypeRadio:
			if len(q.Options) == 0 {
				return Form{}, fmt.Errorf("options are required for radio questions")
			}

			question = RadioQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		case QuestionTypeCheckbox:
			if len(q.Options) == 0 {
				return Form{}, fmt.Errorf("options are required for checkbox questions")
			}

			question = CheckboxQuestion{
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
	form.Questions = questions

	err := s.repo.CreateForm(ctx, form)
	if err != nil {
		return Form{}, err
	}

	return form, nil
}

// func (s *Service) GetFormBase(ctx context.Context, id string) (FormBase, error) {
// 	return s.repo.getFormBase(ctx, id)
//}

func (s *Service) GetForm(ctx context.Context, id string) (Form, error) {
	return s.repo.GetForm(ctx, id)
}

type ListFormsParams struct {
}

func (s *Service) ListForms(ctx context.Context, params ListFormsParams) ([]Form, error) {
	return s.repo.ListForms(ctx, params)
}
