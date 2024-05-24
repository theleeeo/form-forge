package runner

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	formv1 "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/api-go/form/v1/formconnect"
	"github.com/theleeeo/form-forge/app"
	"github.com/theleeeo/form-forge/entrypoints"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/response"
)

type Runner struct {
}

func Run(cfg *Config) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//
	// Create the repository
	//
	formRepo, err := form.NewMySql(&form.MySqlConfig{
		Address:  cfg.RepoCfg.Address,
		User:     cfg.RepoCfg.User,
		Password: cfg.RepoCfg.Password,
		Database: cfg.RepoCfg.Database,
	}, nil)
	if err != nil {
		return err
	}
	defer formRepo.Close()

	resopnseRepo, err := response.NewMySql(&response.MySqlConfig{
		Address:  cfg.RepoCfg.Address,
		User:     cfg.RepoCfg.User,
		Password: cfg.RepoCfg.Password,
		Database: cfg.RepoCfg.Database,
	}, nil)
	if err != nil {
		return err
	}
	defer resopnseRepo.Close()

	// User service
	//
	formSrv := form.NewService(formRepo)
	responseSrv := response.NewService(resopnseRepo)

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
		if err := server.Run(); err != nil {
			log.Printf("error running server: %v", err)
		}
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
