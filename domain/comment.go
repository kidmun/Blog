package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
const (
	CollectionComment = "comments"
)

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID    primitive.ObjectID `bson:"blog_id" json:"blog_id" validate:"required"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Content   string             `bson:"content" json:"content" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
type CommentUsecase interface {
	Create(ctx context.Context, comment *Comment) error
	Update(ctx context.Context,commentID primitive.ObjectID, comment *Comment) error
	Delete(ctx context.Context,commentID primitive.ObjectID) error
	GetAll(ctx context.Context,) ([]*Comment, error)
	GetByBlogID(ctx context.Context,blogID primitive.ObjectID) ([]*Comment, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*Comment, error)
}
type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) error
	Update(ctx context.Context,commentID primitive.ObjectID, comment *Comment) error
	Delete(ctx context.Context,commentID primitive.ObjectID) error
	GetAll(ctx context.Context,) ([]*Comment, error)
	GetByBlogID(ctx context.Context,blogID primitive.ObjectID) ([]*Comment, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*Comment, error)
}