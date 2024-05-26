package form

import (
	"fmt"
	"time"
)

type FormBase struct {
	Id        string
	VersionId string
	Version   uint32

	Title       string
	Description string
	CreatedAt   time.Time
}

type Form struct {
	FormBase
	Questions []Question
}

func constructForm(params CreateFormParams) (Form, error) {
	if params.Title == "" {
		return Form{}, fmt.Errorf("%w: title is required", ErrBadArgs)
	}

	if len(params.Questions) == 0 {
		return Form{}, fmt.Errorf("%w: questions are required", ErrBadArgs)
	}

	id := UUIDNew().String()
	versionId := UUIDNew().String()
	form := Form{
		FormBase: FormBase{
			Id:          id,
			VersionId:   versionId,
			Version:     1,
			Title:       params.Title,
			Description: params.Description,
			CreatedAt:   TimeNow().UTC(),
		},
	}

	questions := make([]Question, 0, len(params.Questions))
	for _, q := range params.Questions {
		var question Question

		base := QuestionBase{
			Title: q.Title,
		}

		switch q.Type {
		case QuestionTypeText:
			question = TextQuestion{
				QuestionBase: base,
			}
		case QuestionTypeRadio:
			if len(q.Options) == 0 {
				return Form{}, fmt.Errorf("%w: options are required for radio questions", ErrBadArgs)
			}

			question = RadioQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		case QuestionTypeCheckbox:
			if len(q.Options) == 0 {
				return Form{}, fmt.Errorf("%w: options are required for checkbox questions", ErrBadArgs)
			}

			question = CheckboxQuestion{
				QuestionBase: base,
				Options:      q.Options,
			}
		default:
			return Form{}, fmt.Errorf("%w: invalid question type: %d", ErrBadArgs, q.Type)
		}

		questions = append(questions, question)

		if err := question.Validate(); err != nil {
			return Form{}, fmt.Errorf("%w: question validation failed: %w", ErrBadArgs, err)
		}
	}
	form.Questions = questions

	return form, nil
}
