package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"iaso/config"
)

func Run(router *gin.Engine, config *config.Config, logger *zap.SugaredLogger) {
	serverLogger := logger.Named("Server")
	srv := &http.Server{
		Addr:    config.HTTPServer.ListenAddr,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverLogger.Infof("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	serverLogger.Info("Shutdown Iaso API Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), config.HTTPServer.GraceShutdownPeriod.Duration)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		serverLogger.Info("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		serverLogger.Infof("timeout of %d seconds.", config.HTTPServer.GraceShutdownPeriod)
	}
	serverLogger.Info("Server exiting")
}

