package runner

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	formv1 "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/api-go/form/v1/formconnect"
	"github.com/theleeeo/form-forge/app"
	"github.com/theleeeo/form-forge/entrypoints"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/response"
)

type Runner struct {
}

func Run(cfg Config) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//
	// PostgreSQL
	//
	dbpool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.RepoCfg.User, cfg.RepoCfg.Password, cfg.RepoCfg.Host, cfg.RepoCfg.Port, cfg.RepoCfg.Database))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	//
	// Repositories
	//
	formRepoPg := form.NewPgRepo(dbpool)
	resopnseRepoPg := response.NewPgRepo(dbpool)

	//
	// User service
	//
	formSrv := form.NewService(formRepoPg)
	responseSrv := response.NewService(resopnseRepoPg)

	//
	// App
	//
	appImpl := app.New(formSrv, responseSrv)

	formGrpcServer := entrypoints.NewFormGRPCServer(appImpl)

	//
	// API Server
	//
	server := entrypoints.NewServer(ctx, &entrypoints.Config{
		GrpcAddr: cfg.GrpcAddress,
		HttpAddr: cfg.HttpAddress,
	})
	server.RegisterService(&formv1.FormService_ServiceDesc, formGrpcServer)
	server.Handle(formconnect.NewFormServiceHandler(entrypoints.NewFormConnectServer(formGrpcServer)))

	httpHandler := entrypoints.NewRestHandler(appImpl)
	httpHandler.RegisterRoutes(server.Mux())
	//
	// Run the server
	//

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting server")
		if err := server.Run(); err != nil {
			log.Printf("error running server: %v", err)
		}
		log.Println("Server stopped")
	}()

	select {
	case <-signalChan:
		log.Println("Received signal, shutting down")
		cancel()

		wg.Wait()

		return nil

	case <-ctx.Done():
		return nil
	}
}
