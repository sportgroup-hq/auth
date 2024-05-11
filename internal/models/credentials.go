package models

import "github.com/google/uuid"

type UserCredential struct {
	UserID       uuid.UUID `json:"-"`
	PasswordHash string    `json:"-"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"password"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	//Picture string `json:"picture"`
}
