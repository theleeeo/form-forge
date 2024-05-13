package runner

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	form_api "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/app"
	"github.com/theleeeo/form-forge/entrypoints"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/repo"
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
	repo, err := repo.NewMySql(&cfg.RepoCfg, nil)
	if err != nil {
		return err
	}
	defer repo.Close()

	// User service
	//
	formSrv := form.NewService(repo)

	//
	// App
	//
	appImpl := app.New(formSrv)

	formGrpcServer := entrypoints.NewFormGRPCHandler(appImpl)

	//
	// GRPC server
	//
	server := entrypoints.NewServer(ctx, &entrypoints.Config{
		GrpcAddr: cfg.GrpcAddress,
		HttpAddr: cfg.HttpAddress,
	})
	server.RegisterService(&form_api.FormService_ServiceDesc, formGrpcServer, form_api.RegisterFormServiceHandlerFromEndpoint)

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
