package httpserver

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func (s *Server) refreshToken(ctx *gin.Context) {
	var reqBody RefreshTokenRequest

	if err := ctx.MustBindWith(&reqBody, binding.JSON); err != nil {
		return
	}

	token, err := s.authService.RefreshToken(ctx, reqBody.RefreshToken)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to refresh token: %s", err)
		ctx.String(500, "Failed to refresh token")
		return
	}

	ctx.JSON(200, gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
		"tokenType":    token.TokenType,
		"expiry":       token.Expiry,
	})
}
