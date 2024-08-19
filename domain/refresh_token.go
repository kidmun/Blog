package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRefreshToken = "refresh_tokens"
)

type RefreshToken struct {
    ID        primitive.ObjectID   `bson:"_id,omitempty"`
    Token     string    `bson:"token"`
    UserID    primitive.ObjectID     `bson:"user_id"`
    ExpiresAt time.Time `bson:"expires_at"`
    Revoked   bool      `bson:"revoked"` 
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id primitive.ObjectID) (*User, error)
	CreateAccessToken(user *User,userId string,  secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
	StoreRefreshToken(ctx context.Context,token *RefreshToken) error
	GetStoredRefreshToken(ctx context.Context, userID primitive.ObjectID) (*RefreshToken, error)
	DeleteTokensByUserID(ctx context.Context, userID primitive.ObjectID) error
	
}

type RefreshTokenRepository interface {
	StoreRefreshToken(ctx context.Context,token *RefreshToken) error
	GetStoredRefreshToken(ctx context.Context, userID primitive.ObjectID) (*RefreshToken, error)
	DeleteTokensByUserID(ctx context.Context, userID primitive.ObjectID) error
	
} 