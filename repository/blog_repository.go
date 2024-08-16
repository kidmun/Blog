package repository

import (
	"Blog/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type taskRepository struct {
	database *mongo.Database
	collection string
}

func NewBlogRepository(database *mongo.Database, collection string) domain.BlogRepository {
	return &taskRepository{
		database: database,
		collection: collection,
	}
}

func (t *taskRepository) Create(ctx context.Context, blog *domain.Blog) (primitive.ObjectID, error) {
	result, err := t.database.Collection(t.collection).InsertOne(ctx, blog)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (t *taskRepository) GetAll(ctx context.Context)([]*domain.Blog, error){
	var blogs []*domain.Blog
	cursor, err := t.database.Collection(t.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &blogs)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (t *taskRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {
	var blog *domain.Blog
	err := t.database.Collection(t.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (t *taskRepository) Update(ctx context.Context,  id primitive.ObjectID, blog *domain.BlogUpdateRequest) error {
	
	var prevBlog *domain.Blog
	err := t.database.Collection(t.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&prevBlog)
	if err != nil {
		return err
	}

	if blog.Title == nil {
		blog.Title = &prevBlog.Title
	}
	if blog.Content == nil {
		blog.Title = &prevBlog.Content
	}
	if blog.Content == nil {
		blog.Title = &prevBlog.Content
	}
	_, err = t.database.Collection(t.collection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": blog})
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	res, err := t.database.Collection(t.collection).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0{
		return mongo.ErrNoDocuments
	}
	return nil
}
