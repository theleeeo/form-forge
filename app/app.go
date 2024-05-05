package app

import (
	"context"

	"go.leeeo.se/form-forge/form"
)

func New(formService *form.Service) *App {
	return &App{
		formService: formService,
	}
}

type App struct {
	formService *form.Service
}

func (a *App) CreateNewForm(ctx context.Context, params form.CreateFormParams) (form.Form, error) {
	return a.formService.CreateNewForm(ctx, params)
}

func (a *App) GetForm(ctx context.Context, id string) (form.Form, error) {
	return a.formService.GetForm(ctx, id)
}
