package entrypoints

import (
	"context"
	"fmt"

	api_go "go.leeeo.se/form-forge/api-go"
	"go.leeeo.se/form-forge/app"
)

var _ api_go.FormServiceServer = &grpcFormHandler{}

func NewFormGRPCHandler(app *app.App) api_go.FormServiceServer {
	return &grpcFormHandler{
		app: app,
	}
}

type grpcFormHandler struct {
	app *app.App

	api_go.UnimplementedFormServiceServer
}

func (g *grpcFormHandler) Create(ctx context.Context, params *api_go.CreateParameters) (*api_go.Form, error) {
	p, err := convertCreateFormParams(params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse parameters: %w", err)
	}

	resp, err := g.app.CreateNewForm(ctx, p)
	if err != nil {
		return nil, err
	}

	form, err := convertForm(ctx, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert form: %w", err)
	}

	return form, nil
}

func (g *grpcFormHandler) GetByID(ctx context.Context, params *api_go.GetByIdParameters) (*api_go.Form, error) {
	f, err := g.app.GetForm(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	form, err := convertForm(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("failed to convert form: %w", err)
	}

	return form, nil
}
