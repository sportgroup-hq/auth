package httpserver

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sportgroup-hq/common-lib/validation"
)

func (s *HTTPServer) Start() error {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.Register(v)
	}

	authPath := r.Group("/auth")

	authPath.GET("/oauth2callback", s.oauthCallback)

	authPath.POST("/login", s.login)
	authPath.POST("/register", s.register)

	authPath.POST("/refresh-token", s.refreshToken)

	slog.Info("Starting HTTP server on " + s.cfg.HTTP.Address + "...")

	return r.Run(s.cfg.HTTP.Address)
}
