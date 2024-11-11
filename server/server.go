package server

import (
	"net/http"

	"github.com/alexkazantsev/go-templ-api/modules/config"
	"github.com/alexkazantsev/go-templ-api/server/middlewares"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"go.uber.org/zap"
)

type Server struct {
	*http.Server
	*gin.Engine

	V1 *gin.RouterGroup
}

func NewServer(cfg config.AppConfig, logger *zap.Logger) *Server {
	engine := gin.New()

	if cfg.Environment.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(cors.AllowAll())
	engine.Use(gin.Recovery())
	engine.Use(middlewares.Logger(logger))

	group := engine.Group("/api/v1")

	return &Server{
		Server: &http.Server{
			Addr:    cfg.GetAddr(),
			Handler: engine,
		},
		Engine: engine,
		V1:     group,
	}
}
