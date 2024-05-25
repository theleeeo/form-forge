package app

import (
	"context"
	"errors"

	"github.com/theleeeo/form-forge/form"
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

func (a *App) ListForms(ctx context.Context, params form.ListFormsParams) ([]form.Form, error) {
	f, err := a.formService.ListForms(ctx, params)
	if err != nil {
		return nil, err
	}

	return f, nil
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

	r, err := a.responseService.ParseResponse(a.convertToFormData(f), resp)
	if err != nil {
		return err
	}

	err = a.responseService.SaveResponse(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) convertToFormData(f form.Form) response.FormData {
	formData := response.FormData{
		Id:        f.Id,
		VersionId: f.VersionId,
	}

	for i, q := range f.Questions {
		var questionType form.QuestionType
		var optionCount int
		switch q := q.(type) {
		case form.TextQuestion:
			questionType = form.QuestionTypeText
			optionCount = 0
		case form.RadioQuestion:
			questionType = form.QuestionTypeRadio
			optionCount = len(q.Options)
		case form.CheckboxQuestion:
			questionType = form.QuestionTypeCheckbox
			optionCount = len(q.Options)
		}

		formData.Questions = append(formData.Questions, response.QuestionData{
			Type:        questionType,
			Order:       i,
			OptionCount: optionCount,
		})
	}

	return formData
}

func (a *App) UpdateForm(ctx context.Context, params form.UpdateFormParams) (form.Form, error) {
	f, err := a.formService.UpdateForm(ctx, params)
	if err != nil {
		if errors.Is(err, form.ErrNotFound) {
			return form.Form{}, ErrFormNotFound
		}
		return form.Form{}, err
	}

	return f, nil
}
