package app

import (
	"context"
	"errors"

	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/models"
	"github.com/theleeeo/form-forge/response"
	"github.com/theleeeo/form-forge/templater"
)

var (
	ErrFormNotFound = errors.New("form not found")
)

func New(formService *form.Service, responseService *response.Service) *App {
	return &App{
		formService:     formService,
		responseService: responseService,
		templater:       templater.New(),
	}
}

type App struct {
	formService     *form.Service
	responseService *response.Service
	templater       *templater.Templater
}

func (a *App) CreateNewForm(ctx context.Context, params form.CreateFormParams) (form.Form, error) {
	return a.formService.CreateNewForm(ctx, params)
}

func (a *App) GetForm(ctx context.Context, id string) (form.Form, error) {
	f, err := a.formService.GetForm(ctx, id)
	if err != nil {
		if errors.Is(err, form.ErrNotFound) {
			return form.Form{}, ErrFormNotFound
		}

		return form.Form{}, err
	}

	return f, nil
}

func (a *App) TemplateForm(ctx context.Context, id string) ([]byte, error) {
	f, err := a.GetForm(ctx, id)
	if err != nil {
		return nil, err
	}

	tpl, err := a.templater.Generate(ctx, f)
	if err != nil {
		return nil, err
	}

	return tpl, nil
}

func (a *App) SubmitResponse(ctx context.Context, formId string, resp map[string][]string) error {
	f, err := a.GetForm(ctx, formId)
	if err != nil {
		return err
	}

	formData, err := a.convertToFormData(ctx, f)
	if err != nil {
		return err
	}

	r, err := a.responseService.ParseResponse(formData, resp)
	if err != nil {
		return err
	}

	err = a.responseService.SaveResponse(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) convertToFormData(ctx context.Context, f form.Form) (response.FormData, error) {
	questions, err := f.Questions(ctx)
	if err != nil {
		return response.FormData{}, err
	}

	formData := response.FormData{
		Id:      f.ID,
		Version: f.Version,
	}

	for i, q := range questions {
		var questionType models.QuestionType
		var optionCount int
		switch q := q.(type) {
		case models.TextQuestion:
			questionType = models.QuestionTypeText
			optionCount = 0
		case models.RadioQuestion:
			questionType = models.QuestionTypeRadio
			optionCount = len(q.Options)
		case models.CheckboxQuestion:
			questionType = models.QuestionTypeCheckbox
			optionCount = len(q.Options)
		}

		formData.Questions = append(formData.Questions, response.QuestionData{
			Type:        questionType,
			Order:       i,
			OptionCount: optionCount,
		})
	}

	return formData, nil
}
