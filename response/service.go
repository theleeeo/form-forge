package response

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/form"
)

func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo *Repo
}

type FormData struct {
	Id        uuid.UUID
	VersionId uuid.UUID
	Questions []QuestionData
}

type QuestionData struct {
	Id          uuid.UUID
	Order       int
	Type        form.QuestionType
	OptionCount int
}

func (s *Service) ParseResponse(formData FormData, resp map[string][]string) (Response, error) {
	r := Response{
		Id:            uuid.New(),
		FormVersionId: formData.VersionId,
		Answers:       make([]Answer, len(formData.Questions)),
		SubmittedAt:   time.Now().UTC(),
	}

	i := 0
	for q, a := range resp {
		if len(a) == 0 {
			return Response{}, fmt.Errorf("answer %s is empty", q)
		}

		questionId, err := uuid.Parse(q)
		if err != nil {
			return Response{}, fmt.Errorf("answer key %s could not be parsed: %w", q, err)
		}

		base := AnswerBase{
			QuestionId: questionId,
		}

		var question QuestionData
		for _, q := range formData.Questions {
			if q.Id == questionId {
				question = q
				break
			}
		}

		if question.Id == uuid.Nil {
			return Response{}, fmt.Errorf("question %s not found", q)
		}

		var answer Answer
		switch question.Type {
		case form.QuestionTypeText:
			if len(a) > 1 {
				return Response{}, fmt.Errorf("text answer %s has more than one value", q)
			}

			answer = TextAnswer{
				AnswerBase: base,
				Value:      a[0],
			}

		case form.QuestionTypeRadio:
			if len(a) > 1 {
				return Response{}, fmt.Errorf("radio answer %s has more than one value", q)
			}

			value, err := strconv.Atoi(a[0])
			if err != nil {
				return Response{}, fmt.Errorf("answer value %s could not be parsed: %w", a[0], err)
			}

			if value < 0 || value >= question.OptionCount {
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

				if value < 0 || value >= question.OptionCount {
					return Response{}, fmt.Errorf("answer value %s is out of range", a)
				}

				values[i] = value
			}

			answer = CheckboxAnswer{
				AnswerBase: base,
				Values:     values,
			}
		}

		r.Answers[i] = answer
		i++
	}

	return r, nil
}

func (s *Service) SaveResponse(ctx context.Context, resp Response) error {
	return s.repo.SaveResponse(ctx, resp)
}
