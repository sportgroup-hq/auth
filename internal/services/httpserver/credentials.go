package httpserver

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sportgroup-hq/auth/internal/models"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *HTTPServer) login(ctx *gin.Context) {
	var reqBody LoginRequest

	if err := ctx.MustBindWith(&reqBody, binding.JSON); err != nil {
		return
	}

	token, err := s.authService.LoginWithEmail(ctx, reqBody.Email, reqBody.Password)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("failed to login: %s", err))
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "wrong email or password",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"access_token":  token.AccessToken,
		"token_type":    token.TokenType,
		"expires_in":    token.Expiry,
		"refresh_token": token.RefreshToken,
	})
}

func (s *HTTPServer) register(ctx *gin.Context) {
	var reqBody models.RegisterRequest

	if err := ctx.MustBindWith(&reqBody, binding.JSON); err != nil {
		return
	}

	token, err := s.authService.Register(ctx, &reqBody)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("failed to register: %s", err))
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(200, gin.H{
		"access_token":  token.AccessToken,
		"token_type":    token.TokenType,
		"expires_in":    token.Expiry,
		"refresh_token": token.RefreshToken,
	})
}
