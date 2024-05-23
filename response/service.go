package response

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/models"
)

func NewService() *Service {
	return &Service{}
}

type Service struct{}

type FormData struct {
	Id        string
	Questions []QuestionData
}

type QuestionData struct {
	Type        models.QuestionType
	Order       int
	OptionCount int
}

func (s *Service) ParseResponse(formData FormData, resp map[string][]string) (Response, error) {
	r := Response{
		Id:      uuid.NewString(),
		FormId:  formData.Id,
		Answers: make([]Answer, len(resp)),
	}

	for q, a := range resp {
		if len(a) == 0 {
			return Response{}, fmt.Errorf("answer %s is empty", q)
		}

		questionOrder, err := strconv.Atoi(q)
		if err != nil {
			return Response{}, fmt.Errorf("answer key %s could not be parsed: %w", q, err)
		}

		base := AnswerBase{
			QuestionOrder: questionOrder,
		}

		var answer Answer
		switch formData.Questions[questionOrder].Type {
		case models.QuestionTypeText:
			answer = TextAnswer{
				AnswerBase: base,
				Value:      a[0],
			}

		case models.QuestionTypeRadio:
			value, err := strconv.Atoi(a[0])
			if err != nil {
				return Response{}, fmt.Errorf("answer value %s could not be parsed: %w", a[0], err)
			}

			if value < 0 || value >= formData.Questions[questionOrder].OptionCount {
				return Response{}, fmt.Errorf("answer value %s is out of range", a)
			}

			answer = RadioAnswer{
				AnswerBase: base,
				Value:      value,
			}

		case models.QuestionTypeCheckbox:
			values := make([]int, len(a))
			for i, v := range a {
				value, err := strconv.Atoi(v)
				if err != nil {
					return Response{}, fmt.Errorf("answer value %s could not be parsed: %w", v, err)
				}

				if value < 0 || value >= formData.Questions[questionOrder].OptionCount {
					return Response{}, fmt.Errorf("answer value %s is out of range", a)
				}

				values[i] = value
			}

			answer = CheckboxAnswer{
				AnswerBase: base,
				Values:     values,
			}
		}

		r.Answers[questionOrder] = answer
	}

	return r, nil
}

func (s *Service) SaveResponse(ctx context.Context, resp Response) error {
	// Save the response to the database.
	return nil
}
