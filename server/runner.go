package server

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func RunServer(lc fx.Lifecycle, server *Server, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			logger.Info("Starting server")

			go func() {
				if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					logger.Fatal("ListenAndServe", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(ctx); err != nil {
				logger.Error("Shutdown", zap.Error(err))

				return err
			}

			return nil
		},
	})
}
