package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
const (
	CollectionLikeDislike = "like_dislike"
)
type LikeDislike struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BlogID    primitive.ObjectID `bson:"blog_id" json:"blog_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	IsLike    bool               `bson:"is_like" json:"is_like"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
type LikeDislikeUsecase interface {
	AddLike(ctx context.Context, likeDislike *LikeDislike) error
	RemoveLike(ctx context.Context,blogID primitive.ObjectID, likeDislikeID primitive.ObjectID) error
	AddDisLike(ctx context.Context,likeDislike *LikeDislike) error
	RemoveDisLike(ctx context.Context,blogID primitive.ObjectID, likeDislikeID primitive.ObjectID) error
}
type LikeDislikeRepository interface {
	AddLike(ctx context.Context, likeDislike *LikeDislike) error
	RemoveLike(ctx context.Context,blogID primitive.ObjectID, likeDislikeID primitive.ObjectID) error
	AddDisLike(ctx context.Context, likeDislike *LikeDislike) error
	RemoveDisLike(ctx context.Context,blogID primitive.ObjectID,  likeDislikeID primitive.ObjectID) error
}