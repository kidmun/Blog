package usecase

import (
	"Blog/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type likeDisLikeUsecase struct {
	likeDisLikeRepository domain.LikeDislikeRepository
	contextTimeout        time.Duration
}

func NewLikeDislikeUsecase(likeDisLikeRepository domain.LikeDislikeRepository, contextTimout time.Duration) domain.LikeDislikeUsecase {
	return &likeDisLikeUsecase{
		likeDisLikeRepository: likeDisLikeRepository,
		contextTimeout:        contextTimout,
	}
}
func (ldu *likeDisLikeUsecase) AddLike(ctx context.Context, likeDislike *domain.LikeDislike) error {
	ctx, cancel := context.WithTimeout(ctx, ldu.contextTimeout)
	defer cancel()
	return ldu.likeDisLikeRepository.AddLike(ctx, likeDislike)
}
func (ldu *likeDisLikeUsecase) RemoveLike(ctx context.Context, blogID primitive.ObjectID, likeDislikeID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, ldu.contextTimeout)
	defer cancel()
	return ldu.likeDisLikeRepository.RemoveDisLike(ctx, blogID, likeDislikeID)
}
func (ldu *likeDisLikeUsecase) AddDisLike(ctx context.Context, likeDislike *domain.LikeDislike) error {
	ctx, cancel := context.WithTimeout(ctx, ldu.contextTimeout)
	defer cancel()
	return ldu.likeDisLikeRepository.AddDisLike(ctx, likeDislike)
}
func (ldu *likeDisLikeUsecase) RemoveDisLike(ctx context.Context, blogID primitive.ObjectID, likeDislikeID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, ldu.contextTimeout)
	defer cancel()
	return ldu.likeDisLikeRepository.RemoveDisLike(ctx, blogID, likeDislikeID)
}
