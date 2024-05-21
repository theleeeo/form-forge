package entrypoints

import (
	"context"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	mb    = 1024 * 1024
	mb256 = 256 * mb
)

func GrpcServerLoggerInterceptor(logger *log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		start := time.Now()

		h, err := handler(ctx, req)
		if err != nil {
			gerr, _ := status.FromError(err)
			logger.Println(
				"GRPC: error handling request,",
				"request_name:", info.FullMethod,
				"request_duration:", time.Since(start).String(),
				"error_code:", gerr.Code(),
				"error_message:", gerr.Message(),
				"error_details:", gerr.Details(),
			)

			return h, err
		}

		logger.Printf(
			"GRPC: successfully handled GRPC request, request_name: %s, request_duration: %s",
			info.FullMethod,
			time.Since(start).String(),
		)

		return h, err
	}
}

func GrpcPanicRecoveryInterceptor(logger *log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (h interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Printf(
					"panic recovered, request_name: %s, error: %v, stack_trace: %s",
					info.FullMethod,
					r,
					string(debug.Stack()),
				)
				err = status.Errorf(codes.Internal, "panic recovered: %v", r)
			}
		}()

		return handler(ctx, req)
	}
}

func NewServer(ctx context.Context, cfg *Config) *server {
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		GrpcPanicRecoveryInterceptor(log.Default()),
		GrpcServerLoggerInterceptor(log.Default()),
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(mb256),
		grpc.MaxSendMsgSize(mb256),
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
	)

	httpServer := &http.Server{
		Addr:         cfg.HttpAddr,
		Handler:      http.NewServeMux(),
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 8 * time.Second,
	}

	s := &server{
		cfg: cfg,
		ctx: ctx,

		grpcServer: grpcServer,
		httpServer: httpServer,
	}

	reflection.Register(s.grpcServer)

	return s
}

type Config struct {
	GrpcAddr string
	HttpAddr string
}

type server struct {
	cfg *Config
	ctx context.Context

	grpcServer *grpc.Server

	httpServer *http.Server
}

func (s *server) Mux() *http.ServeMux {
	return s.httpServer.Handler.(*http.ServeMux)
}

func (s *server) Run() error {
	listener, err := net.Listen("tcp", s.cfg.GrpcAddr)
	if err != nil {
		return err
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				return
			}
			log.Printf("http server error: %v", err)
		}
	}()

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()

	<-s.ctx.Done()

	s.grpcServer.GracefulStop()
	if err := s.httpServer.Shutdown(s.ctx); err != nil {
		return err
	}

	return nil
}

func (s *server) RegisterService(desc *grpc.ServiceDesc, srv any) {
	s.grpcServer.RegisterService(desc, srv)
}

func (s *server) Handle(pattern string, h http.Handler) {
	s.httpServer.Handler.(*http.ServeMux).Handle(pattern, h)
}
