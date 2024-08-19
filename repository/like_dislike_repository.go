package repository

import (
	"Blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type likeDisLikeRepository struct {
	database   *mongo.Database
	collection string
}

func NewLikeDislikeRepository(database *mongo.Database, collection string) domain.LikeDislikeRepository {
	return &likeDisLikeRepository{
		database:   database,
		collection: collection,
	}
}

func (ldr *likeDisLikeRepository) AddLike(ctx context.Context,likeDislike *domain.LikeDislike) error {
	_, err := ldr.database.Collection(ldr.collection).InsertOne(ctx, likeDislike)
	if err != nil {
		return err
	}
	var blog *domain.Blog
	err = ldr.database.Collection(ldr.collection).FindOne(ctx, bson.M{"_id": likeDislike.BlogID}).Decode(&blog)
	if err != nil {
		return err
	}
	blog.Likes = blog.Likes + 1
	_, err = ldr.database.Collection(ldr.collection).UpdateOne(ctx, bson.M{"_id": likeDislike.BlogID}, bson.M{"$set": blog})
	if err != nil {
		return err
	}
	return nil
}
func (ldr *likeDisLikeRepository) RemoveLike(ctx context.Context, blogID primitive.ObjectID,likeDislikeID primitive.ObjectID) error {
	res, err := ldr.database.Collection(ldr.collection).DeleteOne(ctx, bson.M{"_id":likeDislikeID })
	if err != nil {
		return err
	}
	if res.DeletedCount == 0{
		return mongo.ErrNoDocuments
	}
	var blog *domain.Blog
	err = ldr.database.Collection(ldr.collection).FindOne(ctx, bson.M{"_id":blogID}).Decode(&blog)
	if err != nil {
		return err
	}
	blog.Likes = blog.Likes - 1
	_, err = ldr.database.Collection(ldr.collection).UpdateOne(ctx, bson.M{"_id":blogID}, bson.M{"$set": blog})
	if err != nil {
		return err
	}
	return nil
}
func (ldr *likeDisLikeRepository) AddDisLike(ctx context.Context, likeDislike *domain.LikeDislike) error {
	var blog *domain.Blog
	err := ldr.database.Collection(ldr.collection).FindOne(ctx, bson.M{"_id":likeDislike.BlogID}).Decode(&blog)
	if err != nil {
		return err
	}
	blog.DisLikes = blog.DisLikes + 1
	_, err = ldr.database.Collection(ldr.collection).UpdateOne(ctx, bson.M{"_id": likeDislike.BlogID}, bson.M{"$set": blog})
	if err != nil {
		return err
	}
	return nil
}
func (ldr *likeDisLikeRepository) RemoveDisLike(ctx context.Context,blogID primitive.ObjectID, likeDislikeID primitive.ObjectID) error {
	res, err := ldr.database.Collection(ldr.collection).DeleteOne(ctx, bson.M{"_id": likeDislikeID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0{
		return mongo.ErrNoDocuments
	}
	var blog *domain.Blog
	err = ldr.database.Collection(ldr.collection).FindOne(ctx, bson.M{"_id": blogID}).Decode(&blog)
	if err != nil {
		return err
	}
	blog.DisLikes = blog.DisLikes - 1
	_, err = ldr.database.Collection(ldr.collection).UpdateOne(ctx, bson.M{"_id": blogID}, bson.M{"$set": blog})
	if err != nil {
		return err
	}
	return nil
}



