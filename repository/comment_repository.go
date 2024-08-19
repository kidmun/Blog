package repository

import (
	"Blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepository struct {
	database   *mongo.Database
	collection string
}

func NewCommentRepository(database *mongo.Database, collection string) domain.CommentRepository {
	return &commentRepository{
		database:   database,
		collection: collection,
	}
}

func (cr *commentRepository)Create(ctx context.Context,comment *domain.Comment) error{
	_, err := cr.database.Collection(cr.collection).InsertOne(ctx, comment)
	if err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository)GetAll(ctx context.Context) ([]*domain.Comment, error){
	var comments []*domain.Comment
	cursor, err := cr.database.Collection(cr.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (cr *commentRepository) GetByBlogID(ctx context.Context,blogID primitive.ObjectID) ([]*domain.Comment, error){
	var comments []*domain.Comment
	filter := bson.M{"blog_id": blogID}
	cursor, err := cr.database.Collection(cr.collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (cr *commentRepository) GetByUserID(ctx context.Context,userID primitive.ObjectID) ([]*domain.Comment, error){
	var comments []*domain.Comment
	filter := bson.M{"user_id": userID}
	cursor, err := cr.database.Collection(cr.collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (cr *commentRepository) Update(ctx context.Context, commentID primitive.ObjectID, comment *domain.Comment) error{
	_, err := cr.database.Collection(cr.collection).UpdateOne(ctx, bson.M{"_id": commentID}, bson.M{"$set": comment})
	if err != nil {
		return err
	}
	return nil
}
func (cr *commentRepository) Delete(ctx context.Context,commentID primitive.ObjectID) error{
	res, err := cr.database.Collection(cr.collection).DeleteOne(ctx, bson.M{"_id": commentID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0{
		return mongo.ErrNoDocuments
	}
	return nil
}