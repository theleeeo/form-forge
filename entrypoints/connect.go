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

func (f *FormConnectServer) GetByID(ctx context.Context, req *connect.Request[formv1.GetByIDRequest]) (*connect.Response[formv1.GetByIDResponse], error) {
	resp, err := f.grpcServer.GetByID(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}
