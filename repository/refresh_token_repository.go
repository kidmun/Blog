package repository

import (
	"Blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type refreshTokenRepository struct {
	database   *mongo.Database
	collection string
}

func NewRefreshTokenRepository(database *mongo.Database, collection string) domain.RefreshTokenRepository {
	return &refreshTokenRepository{
		database:   database,
		collection: collection,
	}
}

func (r *refreshTokenRepository) StoreRefreshToken(ctx context.Context, token *domain.RefreshToken) error {
	_, err := r.database.Collection(r.collection).InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (r *refreshTokenRepository) GetStoredRefreshToken(ctx context.Context, userID primitive.ObjectID) (*domain.RefreshToken, error) {
	filter := bson.M{"user_id": userID}
	var refreshToken domain.RefreshToken
	err := r.database.Collection(r.collection).FindOne(ctx, filter).Decode(&refreshToken)
	if err != nil {
		return nil, err 
	}
	return &refreshToken, nil
}

func (r *refreshTokenRepository) DeleteTokensByUserID(ctx context.Context, userID primitive.ObjectID) error {
    _, err := r.database.Collection(r.collection).DeleteMany(ctx, bson.M{"user_id": userID})
    return err
}