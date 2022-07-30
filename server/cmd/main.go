package main

import (
	"context"
	"eshop-mock-api/configs"
	"eshop-mock-api/internal/core"
	"eshop-mock-api/internal/routes"
	"eshop-mock-api/internal/services"
	"eshop-mock-api/internal/store"
	"eshop-mock-api/pkg/logger"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configs.SetupConfig()
	accessApiLogger, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", "oss-api-go", "ENV")),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.Get().LogPath+"/access.log"),
	)
	if err != nil {
		panic(err)
	}

	panicApiLogger, err := logger.NewJSONLogger(
		logger.WithField("domain", fmt.Sprintf("%s[%s]", "go-hex-test", "TEST")),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileP(configs.Get().LogPath+"/panic.log"),
	)
	if err != nil {
		panic(err)
	}

	dbConn, err := initDatabase()
	if err != nil {
		panic(err)
	}

	service := services.New(dbConn)
	router := core.New(accessApiLogger, panicApiLogger)
	routes.New(service).RegisterRoutes(router)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + configs.Get().Port,
		Handler:        router,
		ReadTimeout:    time.Duration(configs.Get().APiReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(configs.Get().APiWriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C agin to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

func initDatabase() (db store.Repo, err error) {
	db, err = store.New()
	if err != nil {
		return db, err
	}
	return db, nil
}
