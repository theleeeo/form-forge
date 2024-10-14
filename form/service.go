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

func (s *Service) CreateNewForm(ctx context.Context, params CreateFormParams) (Form, []Question, error) {
	form, questions, err := constructForm(params)
	if err != nil {
		return Form{}, nil, err
	}

	err = s.repo.CreateForm(ctx, form, questions)
	if err != nil {
		return Form{}, nil, err
	}

	return form, questions, nil
}

func (s *Service) GetForm(ctx context.Context, baseId uuid.UUID) (Form, error) {
	if baseId == uuid.Nil {
		return Form{}, fmt.Errorf("baseId is required")
	}

	return s.repo.GetLatestVersionOfBase(ctx, baseId)
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

func (s *Service) UpdateForm(ctx context.Context, params UpdateFormParams) (Form, []Question, error) {
	baseForm, err := s.GetForm(ctx, params.Id)
	if err != nil {
		return Form{}, nil, fmt.Errorf("geting form: %w", err)
	}

	form, questions, err := constructForm(params.CreateFormParams)
	if err != nil {
		return Form{}, nil, err
	}
	form.BaseId = params.Id
	form.Version = baseForm.Version + 1

	if err := s.repo.CreateForm(ctx, form, questions); err != nil {
		return Form{}, nil, fmt.Errorf("creating form: %w", err)
	}

	return form, questions, nil
}

type GetQuestionsParams struct {
	BaseId    uuid.UUID
	VersionId uuid.UUID
}

func (s *Service) GetQuestions(ctx context.Context, params GetQuestionsParams) ([]Question, error) {
	if params.VersionId != uuid.Nil {
		return s.repo.GetQuestionsOfVersion(ctx, params.VersionId)
	} else if params.BaseId != uuid.Nil {
		return s.repo.GetQuestions(ctx, params.BaseId)
	}

	return nil, fmt.Errorf("either baseId or versionId is required")
}
