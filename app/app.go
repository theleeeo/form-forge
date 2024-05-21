package app

import (
	"context"
	"errors"

	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/repo"
	"github.com/theleeeo/form-forge/templater"
)

var (
	ErrFormNotFound = errors.New("form not found")
)

func New(formService *form.Service) *App {
	return &App{
		formService: formService,
		templater:   templater.New(),
	}
}

type App struct {
	formService *form.Service
	templater   *templater.Templater
}

func (a *App) CreateNewForm(ctx context.Context, params form.CreateFormParams) (form.Form, error) {
	return a.formService.CreateNewForm(ctx, params)
}

func (a *App) GetForm(ctx context.Context, id string) (form.Form, error) {
	f, err := a.formService.GetForm(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
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
