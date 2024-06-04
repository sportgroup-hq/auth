package httpserver

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) oauthCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		url := s.authService.GetOAuthConsentURL(ctx)
		ctx.Redirect(http.StatusSeeOther, url)
		return
	}

	token, err := s.authService.ExchangeProvidersCode(ctx, code)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to exchange token: %s", err)
		ctx.String(http.StatusInternalServerError, "Failed to exchange token")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
		"tokenType":    token.TokenType,
		"expiry":       token.Expiry,
	})
}
