package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupRequest struct {
	UserName string `form:"userName" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) (primitive.ObjectID, error)
	GetUserByEmail(c context.Context, email string) (*User, error)
	CreateAccessToken(user *User, userId string, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	UpdateUserVerificationStatus(ctx context.Context, userID primitive.ObjectID) error
}