package entrypoints

import (
	"context"
	"fmt"

	form_api "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/app"
)

var _ form_api.FormServiceServer = &grpcFormHandler{}

func NewFormGRPCHandler(app *app.App) form_api.FormServiceServer {
	return &grpcFormHandler{
		app: app,
	}
}

type grpcFormHandler struct {
	app *app.App

	form_api.UnimplementedFormServiceServer
}

func (g *grpcFormHandler) Create(ctx context.Context, params *form_api.CreateRequest) (*form_api.CreateResponse, error) {
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

	return &form_api.CreateResponse{
		Form: form,
	}, nil
}

func (g *grpcFormHandler) GetByID(ctx context.Context, params *form_api.GetByIDRequest) (*form_api.GetByIDResponse, error) {
	f, err := g.app.GetForm(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	form, err := convertForm(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("failed to convert form: %w", err)
	}

	return &form_api.GetByIDResponse{
		Form: form,
	}, nil
}
