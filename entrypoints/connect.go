package entrypoints

import (
	"context"

	"connectrpc.com/connect"
	formv1 "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/api-go/form/v1/formconnect"
)

var _ formconnect.FormServiceHandler = &FormConnectServer{}

func NewFormConnectServer(grpcServer *formGrpcServer) *FormConnectServer {
	return &FormConnectServer{grpcServer: grpcServer}
}

type FormConnectServer struct {
	grpcServer *formGrpcServer
}

func (f *FormConnectServer) Create(ctx context.Context, req *connect.Request[formv1.CreateRequest]) (*connect.Response[formv1.CreateResponse], error) {
	resp, err := f.grpcServer.Create(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (f *FormConnectServer) GetById(ctx context.Context, req *connect.Request[formv1.GetByIdRequest]) (*connect.Response[formv1.GetByIdResponse], error) {
	resp, err := f.grpcServer.GetById(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (f *FormConnectServer) List(ctx context.Context, req *connect.Request[formv1.ListRequest]) (*connect.Response[formv1.ListResponse], error) {
	resp, err := f.grpcServer.List(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (f *FormConnectServer) Update(ctx context.Context, req *connect.Request[formv1.UpdateRequest]) (*connect.Response[formv1.UpdateResponse], error) {
	resp, err := f.grpcServer.Update(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (f *FormConnectServer) GetQuestions(ctx context.Context, req *connect.Request[formv1.GetQuestionsRequest]) (*connect.Response[formv1.GetQuestionsResponse], error) {
	resp, err := f.grpcServer.GetQuestions(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
