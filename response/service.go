package response

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/form"
)

func NewService(repo *MySqlRepo) *Service {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo *MySqlRepo
}

type FormData struct {
	Id        string
	Version   int
	Questions []QuestionData
}

type QuestionData struct {
	Order       int
	Type        form.QuestionType
	OptionCount int
}

func (s *Service) ParseResponse(formData FormData, resp map[string][]string) (Response, error) {
	r := Response{
		Id:          uuid.NewString(),
		FormId:      formData.Id,
		FormVersion: formData.Version,
		Answers:     make([]Answer, len(resp)),
	}

	for q, a := range resp {
		if len(a) == 0 {
			return Response{}, fmt.Errorf("answer %s is empty", q)
		}

		questionOrder, err := strconv.Atoi(q)
		if err != nil {
			return Response{}, fmt.Errorf("answer key %s could not be parsed: %w", q, err)
		}

		if questionOrder < 0 || questionOrder >= len(formData.Questions) {
			return Response{}, fmt.Errorf("answer key %s is out of range", q)
		}

		base := AnswerBase{
			QuestionOrder: questionOrder,
		}

		var answer Answer
		switch formData.Questions[questionOrder].Type {
		case form.QuestionTypeText:
			answer = TextAnswer{
				AnswerBase: base,
				Value:      a[0],
			}

		case form.QuestionTypeRadio:
			if len(a) > 1 {
				return Response{}, fmt.Errorf("answer %s has more than one value", q)
			}

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

		case form.QuestionTypeCheckbox:
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
	return s.repo.SaveResponse(ctx, resp)
}
