package runner

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
	formv1 "github.com/theleeeo/form-forge/api-go/form/v1"
	"github.com/theleeeo/form-forge/api-go/form/v1/formconnect"
	"github.com/theleeeo/form-forge/app"
	"github.com/theleeeo/form-forge/entrypoints"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/response"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respCatcher := httptest.NewRecorder()
		next.ServeHTTP(respCatcher, r)

		log.Printf("%s %s %d", r.Method, r.URL.Path, respCatcher.Code)

		copyHeaders(w.Header(), respCatcher.Header())
		w.WriteHeader(respCatcher.Code)
		_, _ = w.Write(respCatcher.Body.Bytes())
	})
}

func copyHeaders(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

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
	apiServer := entrypoints.NewServer(ctx, &entrypoints.Config{
		Addr: cfg.ApiAddr,
	})
	apiServer.RegisterService(&formv1.FormService_ServiceDesc, formGrpcServer)

	connectPath, connectHandler := formconnect.NewFormServiceHandler(entrypoints.NewFormConnectServer(formGrpcServer))
	apiServer.Handle(connectPath, cors.AllowAll().Handler(LogMiddleware(connectHandler)))

	//
	// Public Server
	//
	mux := http.NewServeMux()
	publicServer := http.Server{
		Addr:    cfg.PublicAddr,
		Handler: mux,
	}

	httpHandler := entrypoints.NewRestHandler(appImpl)
	httpHandler.RegisterRoutes(mux)
	//
	// Run the server
	//

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting api server")
		if err := apiServer.Run(); err != nil {
			log.Printf("error running api server: %v", err)
		}
		log.Println("Api server stopped")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting public server")
		if err := publicServer.ListenAndServe(); err != nil {
			log.Printf("error running public server: %v", err)
		}
		log.Println("public server stopped")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		log.Println("Shutting down public server")
		if err := publicServer.Shutdown(ctx); err != nil {
			log.Printf("error shutting down public server: %v", err)
		}
		log.Println("Public server stopped")
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
