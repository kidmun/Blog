package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}
type ForgotPassowrdUsecase interface {
	GetUserByEmail(c context.Context, email string) (*User, error)
	UpdatePassword(ctx context.Context, userID primitive.ObjectID, newPassword ResetPasswordRequest) error
}