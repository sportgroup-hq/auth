package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) oauthCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		url := s.authService.GetOAuthConsentURL(ctx)
		ctx.Redirect(http.StatusSeeOther, url)
		return
	}

	token, err := s.authService.ExchangeProvidersCode(ctx, code)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to exchange token: %s", err)
		ctx.String(500, "Failed to exchange token")
		return
	}

	ctx.JSON(200, gin.H{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"token_type":    token.TokenType,
		"expiry":        token.Expiry,
	})
}
