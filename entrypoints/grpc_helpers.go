package entrypoints

import (
	"context"
	"fmt"

	form_api "go.leeeo.se/form-forge/api-go/form/v1"
	"go.leeeo.se/form-forge/form"
	"go.leeeo.se/form-forge/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCreateFormParams(params *form_api.CreateRequest) (form.CreateFormParams, error) {
	qs := make([]form.CreateQuestionParams, len(params.Questions))
	for i, q := range params.Questions {
		if q.Type == form_api.Question_TYPE_TEXT {
			q.Options = nil
		} else if len(q.Options) == 0 {
			return form.CreateFormParams{}, fmt.Errorf("question options must not be empty for question type %s", q.Type.String())
		}

		t, err := convertQuestionType(q.Type)
		if err != nil {
			return form.CreateFormParams{}, err
		}

		qs[i] = form.CreateQuestionParams{
			Type:    t,
			Title:   q.Title,
			Options: q.Options,
		}
	}

	return form.CreateFormParams{
		Title:     params.Title,
		Questions: qs,
	}, nil
}

func convertQuestionType(t form_api.Question_Type) (models.QuestionType, error) {
	switch t {
	case form_api.Question_TYPE_TEXT:
		return models.QuestionTypeText, nil
	case form_api.Question_TYPE_RADIO:
		return models.QuestionTypeRadio, nil
	case form_api.Question_TYPE_CHECKBOX:
		return models.QuestionTypeCheckbox, nil
	default:
		return models.QuestionType(-1), fmt.Errorf("invalid question type: %s", t.String())
	}
}

func convertForm(ctx context.Context, f form.Form) (*form_api.Form, error) {
	qs, err := f.Questions(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get questions: %w", err)
	}

	questions := make([]*form_api.Question, 0, len(qs))
	for _, q := range qs {
		questions = append(questions, convertQuestionToProto(q))
	}

	return &form_api.Form{
		Id:        f.ID,
		Title:     f.Title,
		Questions: questions,
		CreatedAt: timestamppb.New(f.CreatedAt),
	}, nil
}

func convertQuestionToProto(q models.Question) *form_api.Question {
	base := q.Question()

	var questionType form_api.Question_Type

	switch q.(type) {
	case models.TextQuestion:
		questionType = form_api.Question_TYPE_TEXT
	case models.RadioQuestion:
		questionType = form_api.Question_TYPE_RADIO
	case models.CheckboxQuestion:
		questionType = form_api.Question_TYPE_CHECKBOX
	default:
		// This should never happen
		panic(fmt.Sprintf("unhandled question type: %T", q))
	}

	var options []string
	switch q := q.(type) {
	case models.RadioQuestion:
		options = q.Options
	case models.CheckboxQuestion:
		options = q.Options
	}

	qp := &form_api.Question{
		Title:   base.Title,
		Type:    questionType,
		Options: options,
	}
	return qp
}
