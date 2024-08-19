package domain

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionBlog = "blogs"
)

type Blog struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Tags []string `json:"tags" bson:"tags"`
	Likes int `json:"like" bson:"like"`
	DisLikes int `json:"disLike" bson:"disLike"`
	Date time.Time `json:"date" bson:"date"`
}


type CreateBlogRequest struct {
	Title string `json:"title" bson:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" bson:"content" binding:"required,min=3,max=100"`
	Tags []string `json:"tags" bson:"tags"`
}
type BlogUpdateRequest struct {
	Title   *string `json:"title,omitempty" bson:"title,omitempty" binding:"omitempty,min=3,max=100"`
	Content *string `json:"content,omitempty" bson:"content,omitempty" binding:"omitempty,min=3,max=100"`
	Tags    *[]string `json:"tags,omitempty" bson:"tags,omitempty"`
}

type BlogUsecase interface {
	Create(ctx context.Context, blog *Blog) (primitive.ObjectID, error)
	GetAll(ctx context.Context)([]*Blog, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Blog, error)
	Update(ctx context.Context, id primitive.ObjectID, blog *BlogUpdateRequest) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
type BlogRepository interface {
	Create(ctx context.Context, blog *Blog) (primitive.ObjectID, error)
	GetAll(ctx context.Context)([]*Blog, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Blog, error)
	Update(ctx context.Context,id primitive.ObjectID,  blog *BlogUpdateRequest) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	
}