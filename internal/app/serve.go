package app

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Reshnyak/AcornsTask/internal/configs"
	"github.com/Reshnyak/AcornsTask/internal/loggers"
	routers "github.com/Reshnyak/AcornsTask/pkg/routers/gin"
	// "google.golang.org/grpc/profiling/service"
	// "google.golang.org/grpc/profiling/service"
)

func Run() {
	cfg, err := configs.ParseFlags()
	if err != nil {
		log.Printf("Config error. Set default values: %s", err)
	}
	log := loggers.InitLogger(cfg)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// initialize dbs
	// storeTasks, err := storage.NewTaskStore(ctx, &cfg.TaskStore)
	// if err != nil {
	// 	log.Fatal().Any("ERROR store tasks", err)
	// }

	log.Printf("Database connection created successfully")
	router := routers.InitRouter(cfg)
	httpServer := http.Server{
		Addr:         cfg.HTTPServer.Address,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		Handler:      router,
	}
	//
	//	authGRPCHandlers := auth.NewAuthHandlers(&storage, passwordHasher, jwtManager, buildinfo.New())

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal().Any("ListenAndServe", err)
		}
	}()
	log.Info().Any("server listening:port", cfg.HTTPServer.Address)

	<-ctx.Done()

	closeCtx, _ := context.WithTimeout(context.Background(), time.Second*5)
	if err := httpServer.Shutdown(closeCtx); err != nil {
		log.Error().Any("httpServer.Shutdown", err)
	}
}
