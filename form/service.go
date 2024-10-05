package form

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	// These variables are used to make the code testable
	UUIDNew = uuid.New
	TimeNow = time.Now
)

var (
	ErrNotFound = errors.New("not found")
	ErrBadArgs  = errors.New("bad arguments")
)

func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

type Service struct {
	repo *Repo
}

type CreateFormParams struct {
	Title       string
	Description string
	Questions   []CreateQuestionParams
}

type CreateQuestionParams struct {
	Type  QuestionType
	Title string
	// Options is only required for radio and checkbox questions
	Options []string
}

func (s *Service) CreateNewForm(ctx context.Context, params CreateFormParams) (Form, error) {
	form, err := constructForm(params)
	if err != nil {
		return Form{}, err
	}

	err = s.repo.CreateForm(ctx, form)
	if err != nil {
		return Form{}, err
	}

	return form, nil
}

func (s *Service) GetForm(ctx context.Context, id uuid.UUID) (Form, error) {
	if id == uuid.Nil {
		return Form{}, fmt.Errorf("id is required")
	}

	return s.repo.GetLatestVersionOfForm(ctx, id)
}

type ListFormsParams struct {
}

func (s *Service) ListForms(ctx context.Context, params ListFormsParams) ([]Form, error) {
	return s.repo.ListForms(ctx, params)
}

type UpdateFormParams struct {
	Id uuid.UUID
	CreateFormParams
}

func (s *Service) UpdateForm(ctx context.Context, params UpdateFormParams) (Form, error) {
	baseForm, err := s.GetForm(ctx, params.Id)
	if err != nil {
		return Form{}, fmt.Errorf("failed to get form: %w", err)
	}

	form, err := constructForm(params.CreateFormParams)
	if err != nil {
		return Form{}, err
	}
	form.BaseId = params.Id
	form.Version = baseForm.Version + 1

	return form, s.repo.CreateForm(ctx, form)
}
