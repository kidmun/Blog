package usecase

import (
	"Blog/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
	contextTimeout    time.Duration
}

func NewCommentUsecase(commentRepository domain.CommentRepository, contextTimout time.Duration) domain.CommentRepository {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeout: contextTimout,
	}
}
func (cu *commentUsecase)Create(ctx context.Context,comment *domain.Comment) error{
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.Create(ctx, comment)
}

func (cu *commentUsecase)GetAll(ctx context.Context) ([]*domain.Comment, error){
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.GetAll(ctx)
}
func (cu *commentUsecase) GetByBlogID(ctx context.Context,blogID primitive.ObjectID) ([]*domain.Comment, error){
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.GetByBlogID(ctx, blogID)
	
}
func (cu *commentUsecase) GetByUserID(ctx context.Context,userID primitive.ObjectID) ([]*domain.Comment, error){
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.GetByUserID(ctx, userID)
}
func (cu *commentUsecase) Update(ctx context.Context, commentID primitive.ObjectID, comment *domain.Comment) error{
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.Update(ctx, commentID, comment)
}
func (cu *commentUsecase) Delete(ctx context.Context,commentID primitive.ObjectID) error{
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.commentRepository.Delete(ctx, commentID)
}