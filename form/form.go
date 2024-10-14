package form

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Form struct {
	BaseId    uuid.UUID
	VersionId uuid.UUID
	Version   uint32

	Title       string
	Description string
	CreatedAt   time.Time
}

func constructForm(params CreateFormParams) (Form, []Question, error) {
	if params.Title == "" {
		return Form{}, nil, fmt.Errorf("%w: title is required", ErrBadArgs)
	}

	if len(params.Questions) == 0 {
		return Form{}, nil, fmt.Errorf("%w: questions are required", ErrBadArgs)
	}

	form := Form{
		BaseId:      UUIDNew(),
		VersionId:   UUIDNew(),
		Version:     1,
		Title:       params.Title,
		Description: params.Description,
		CreatedAt:   TimeNow().UTC(),
	}

	questions := make([]Question, 0, len(params.Questions))
	for _, q := range params.Questions {
		var question Question

		base := QuestionBase{
			Id:    UUIDNew(),
			Title: q.Title,
		}

		switch q.Type {
		case QuestionTypeText:
			question = TextQuestion{
				QuestionBase: base,
			}
		case QuestionTypeRadio:
			if len(q.Options) == 0 {
				return Form{}, nil, fmt.Errorf("%w: options are required for radio questions", ErrBadArgs)
			}

			question = RadioQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		case QuestionTypeCheckbox:
			if len(q.Options) == 0 {
				return Form{}, nil, fmt.Errorf("%w: options are required for checkbox questions", ErrBadArgs)
			}

			question = CheckboxQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		default:
			return Form{}, nil, fmt.Errorf("%w: invalid question type: %d", ErrBadArgs, q.Type)
		}

		questions = append(questions, question)

		if err := question.Validate(); err != nil {
			return Form{}, nil, fmt.Errorf("%w: question validation failed: %w", ErrBadArgs, err)
		}
	}

	return form, questions, nil
}
