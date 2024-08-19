package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogoutUsecase interface {
	DeleteTokensByUserID(ctx context.Context, userID primitive.ObjectID) error
}