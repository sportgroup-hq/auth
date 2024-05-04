package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sportgroup-hq/auth/internal/bootstrap"
)

func main() {
	deps, err := bootstrap.Up()
	if err != nil {
		panic(err)
	}

	//slog.SetDefault(slog.New(logger.NewLogger(os.Stdout, config.Get().Log.Level)))

	user, err := deps.UserService.GetUserByID(context.Background(), uuid.New())
	if err != nil {
		panic(err)
	}

	fmt.Println(user.Id)

	//slog.Info("Starting auth server on " + deps.Config.Address)

	//conf := &oauth2.Config{
	//	ClientID:     cfg.Oauth.Google.ClientID,
	//	ClientSecret: cfg.Oauth.Google.ClientSecret,
	//	RedirectURL:  cfg.Oauth.Google.RedirectURL,
	//	Scopes:       cfg.Oauth.Google.Scopes,
	//	Endpoint:     google.Endpoint,
	//}
	//r.GET("/oauth2callback", func(c *gin.Context) {
	//	code := c.Query("code")
	//	if code == "" {
	//		url := conf.AuthCodeURL(randString(64), oauth2.AccessTypeOnline)
	//		c.Redirect(302, url)
	//	}
	//
	//	token, err := conf.Exchange(c, code)
	//	if err != nil {
	//		c.String(500, "Failed to exchange token: %s", err)
	//		return
	//	}
	//
	//	idTokenStr := token.Extra("id_token").(string)
	//
	//	idTokenPayload, err := idtoken.ParsePayload(idTokenStr)
	//	if err != nil {
	//		c.String(500, "Failed to parse ID token: %s", err)
	//		return
	//	}
	//
	//	c.JSON(200, gin.H{
	//		"email":       idTokenPayload.Claims["email"],
	//		"name":        idTokenPayload.Claims["name"],
	//		"given_name":  idTokenPayload.Claims["given_name"],
	//		"family_name": idTokenPayload.Claims["family_name"],
	//		"picture":     idTokenPayload.Claims["picture"],
	//		"token":       token.AccessToken,
	//	})
	//})

	//if err = deps.HTTPServer.Start(); err != nil {
	//	panic(err)
	//}
}
