package repository

import (
	"Blog/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(database *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) (primitive.ObjectID, error) {
	collection := ur.database.Collection(ur.collection)
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (ur *userRepository) Fetch(ctx context.Context)([]*domain.User, error){
	collection := ur.database.Collection(ur.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	var users []*domain.User

	err = cursor.All(ctx, &users)
	if users == nil {
		return nil, err
	}
	return users, err
}

func (ur *userRepository)GetByEmail(c context.Context, email string) (*domain.User, error){
	collection := ur.database.Collection(ur.collection)
	var user *domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	
	return user, err
}
func (ur *userRepository) GetByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user *domain.User
	err := collection.FindOne(c, bson.M{"_id": id}).Decode(&user)
	return user, err
}
func (ur *userRepository) UpdateUserVerificationStatus(ctx context.Context, userID primitive.ObjectID) error {


    filter := bson.D{primitive.E{Key: "_id", Value: userID}}
    update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "verified", Value: true}}}}
    result, err := ur.database.Collection(ur.collection).UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }
    if result.MatchedCount == 0 {
        return errors.New("user not found")
    }
    return nil
}



