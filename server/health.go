package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupHealth(s *Server, db *sql.DB) {
	s.Engine.GET("/health", func(ctx *gin.Context) {
		if err := db.Ping(); err != nil {
			ctx.String(http.StatusInternalServerError, "Database is not available")

			return
		}

		ctx.String(http.StatusOK, "OK")
	})
}
