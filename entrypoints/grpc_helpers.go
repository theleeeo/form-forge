package entrypoints

import (
	"fmt"

	"github.com/google/uuid"
	form_api "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/form"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUpdateFormParams(params *form_api.UpdateRequest) (form.UpdateFormParams, error) {
	baseUuid, err := uuid.Parse(params.BaseId)
	if err != nil {
		return form.UpdateFormParams{}, fmt.Errorf("could not parse id: %w", err)
	}

	return form.UpdateFormParams{
		Id:               baseUuid,
		CreateFormParams: convertCreateFormParams(params.NewForm),
	}, nil
}

func convertCreateFormParams(params *form_api.CreateRequest) form.CreateFormParams {
	qs := make([]form.CreateQuestionParams, len(params.Questions))
	for i, q := range params.Questions {
		qs[i] = convertCreateQuestionParams(q)
	}

	return form.CreateFormParams{
		Title:       params.Title,
		Description: params.Description,
		Questions:   qs,
	}
}

func convertCreateQuestionParams(qp *form_api.CreateQuestionParameters) form.CreateQuestionParams {
	switch q := qp.Question.(type) {
	case *form_api.CreateQuestionParameters_Text:
		return form.CreateQuestionParams{
			Type:  form.QuestionTypeText,
			Title: q.Text.Title,
		}
	case *form_api.CreateQuestionParameters_Radio:
		return form.CreateQuestionParams{
			Type:    form.QuestionTypeRadio,
			Title:   q.Radio.Title,
			Options: q.Radio.Options,
		}
	case *form_api.CreateQuestionParameters_Checkbox:
		return form.CreateQuestionParams{
			Type:    form.QuestionTypeCheckbox,
			Title:   q.Checkbox.Title,
			Options: q.Checkbox.Options,
		}
	default:
		// This should never happen
		panic(fmt.Sprintf("unhandled question type: %T", q))
	}
}

func convertForm(f form.Form) *form_api.Form {
	return &form_api.Form{
		BaseId:      f.BaseId.String(),
		VersionId:   f.VersionId.String(),
		Version:     f.Version,
		Title:       f.Title,
		Description: f.Description,
		CreatedAt:   timestamppb.New(f.CreatedAt),
	}
}

func convertQuestion(q form.Question) *form_api.Question {
	switch q := q.(type) {
	case form.TextQuestion:
		return &form_api.Question{
			Question: &form_api.Question_Text{
				Text: &form_api.TextQuestion{
					Title: q.Question().Title,
				},
			},
		}

	case form.RadioQuestion:
		return &form_api.Question{
			Question: &form_api.Question_Radio{
				Radio: &form_api.RadioQuestion{
					Title:   q.Question().Title,
					Options: q.Options,
				},
			},
		}

	case form.CheckboxQuestion:
		return &form_api.Question{
			Question: &form_api.Question_Checkbox{
				Checkbox: &form_api.CheckboxQuestion{
					Title:   q.Question().Title,
					Options: q.Options,
				},
			},
		}

	default:
		// This should never happen
		panic(fmt.Sprintf("unhandled question type: %T", q))
	}
}
