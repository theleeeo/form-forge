package entrypoints

import (
	"context"
	"errors"

	form_api "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/app"
	"github.com/theleeeo/form-forge/form"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	resp, err := g.app.CreateNewForm(ctx, convertCreateFormParams(params))
	if err != nil {
		return nil, err
	}

	return &form_api.CreateResponse{
		Form: convertForm(resp),
	}, nil
}

func (g *formGrpcServer) GetById(ctx context.Context, params *form_api.GetByIdRequest) (*form_api.GetByIdResponse, error) {
	f, err := g.app.GetForm(ctx, params.Id)
	if err != nil {
		return nil, err
	}

	return &form_api.GetByIdResponse{
		Form: convertForm(f),
	}, nil
}

func (g *formGrpcServer) List(ctx context.Context, params *form_api.ListRequest) (*form_api.ListResponse, error) {
	f, err := g.app.ListForms(ctx, form.ListFormsParams{})
	if err != nil {
		return nil, err
	}

	var forms []*form_api.Form
	for _, form := range f {
		forms = append(forms, convertForm(form))
	}

	return &form_api.ListResponse{
		Forms: forms,
		Pagination: &form_api.ResponsePagination{
			Total: uint64(len(forms)),
		},
	}, nil
}

func (g *formGrpcServer) Update(ctx context.Context, params *form_api.UpdateRequest) (*form_api.UpdateResponse, error) {
	resp, err := g.app.UpdateForm(ctx, convertUpdateFormParams(params))
	if err != nil {
		if errors.Is(err, app.ErrFormNotFound) {
			return nil, status.Errorf(codes.NotFound, "form not found")
		}

		return nil, err
	}

	return &form_api.UpdateResponse{
		Form: convertForm(resp),
	}, nil
}
