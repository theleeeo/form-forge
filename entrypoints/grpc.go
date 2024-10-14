package entrypoints

import (
	"context"
	"errors"

	"github.com/google/uuid"
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
}

func (g *formGrpcServer) Create(ctx context.Context, params *form_api.CreateRequest) (*form_api.CreateResponse, error) {
	resp, _, err := g.app.CreateNewForm(ctx, convertCreateFormParams(params))
	if err != nil {
		return nil, err
	}

	return &form_api.CreateResponse{
		BaseId:    resp.BaseId.String(),
		VersionId: resp.VersionId.String(),
	}, nil
}

func (g *formGrpcServer) GetById(ctx context.Context, params *form_api.GetByIdRequest) (*form_api.GetByIdResponse, error) {
	baseUUID, err := uuid.Parse(params.BaseId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not parse base_id: %v", err)
	}

	f, err := g.app.GetForm(ctx, baseUUID)
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
	p, err := convertUpdateFormParams(params)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not convert update params: %v", err)
	}

	resp, _, err := g.app.UpdateForm(ctx, p)
	if err != nil {
		if errors.Is(err, app.ErrFormNotFound) {
			return nil, status.Errorf(codes.NotFound, "form not found")
		}

		return nil, err
	}

	return &form_api.UpdateResponse{
		BaseId:    resp.BaseId.String(),
		VersionId: resp.VersionId.String(),
	}, nil
}

func (g *formGrpcServer) GetQuestions(ctx context.Context, params *form_api.GetQuestionsRequest) (*form_api.GetQuestionsResponse, error) {
	var baseUUID, versionUUID uuid.UUID
	var err error

	if params.BaseId != "" {
		baseUUID, err = uuid.Parse(params.BaseId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "could not parse base_id: %v", err)
		}
	}

	if params.VersionId == "" {
		if baseUUID == uuid.Nil {
			return nil, status.Errorf(codes.InvalidArgument, "base_id is required")
		}
	}

	q, err := g.app.GetQuestions(ctx, form.GetQuestionsParams{
		BaseId:    baseUUID,
		VersionId: versionUUID,
	})
	if err != nil {
		return nil, err
	}

	var questions []*form_api.Question
	for _, question := range q {
		questions = append(questions, convertQuestion(question))
	}

	return &form_api.GetQuestionsResponse{
		Questions: questions,
	}, nil
}
