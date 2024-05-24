package entrypoints

import (
	"context"
	"fmt"

	form_api "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/app"
	"github.com/theleeeo/form-forge/form"
)

var _ form_api.FormServiceServer = &formGrpcServer{}

func NewFormGRPCServer(app *app.App) *formGrpcServer {
	return &formGrpcServer{
		app: app,
	}
}

type formGrpcServer struct {
	app *app.App

	form_api.UnimplementedFormServiceServer
}

func (g *formGrpcServer) Create(ctx context.Context, params *form_api.CreateRequest) (*form_api.CreateResponse, error) {
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

func (g *formGrpcServer) GetById(ctx context.Context, params *form_api.GetByIdRequest) (*form_api.GetByIdResponse, error) {
	f, err := g.app.GetForm(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	form, err := convertForm(ctx, f)
	if err != nil {
		return nil, fmt.Errorf("failed to convert form: %w", err)
	}

	return &form_api.GetByIdResponse{
		Form: form,
	}, nil
}

func (g *formGrpcServer) List(ctx context.Context, params *form_api.ListRequest) (*form_api.ListResponse, error) {
	f, err := g.app.ListForms(ctx, form.ListFormsParams{})
	if err != nil {
		return nil, err
	}

	var forms []*form_api.Form
	for _, form := range f {
		f, err := convertForm(ctx, form)
		if err != nil {
			return nil, fmt.Errorf("failed to convert form: %w", err)
		}

		forms = append(forms, f)
	}

	return &form_api.ListResponse{
		Forms: forms,
		Pagination: &form_api.ResponsePagination{
			Total: uint64(len(forms)),
		},
	}, nil
}
