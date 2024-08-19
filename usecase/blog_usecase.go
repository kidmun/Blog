package usecase

import (
	"Blog/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type blogUsecase struct {
	blogRepository domain.BlogRepository
	contextTimeout time.Duration
}

func NewBlogUsecase(blogRepository domain.BlogRepository, contextTimout time.Duration) domain.BlogUsecase {
	return &blogUsecase{
		blogRepository: blogRepository,
		contextTimeout: contextTimout,
	}
}
func (b *blogUsecase) Create(ctx context.Context, blog *domain.Blog) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.blogRepository.Create(ctx, blog)
}
func (b *blogUsecase) GetAll(ctx context.Context)([]*domain.Blog, error) {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.blogRepository.GetAll(ctx)
}

func (b *blogUsecase) GetByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {	
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.blogRepository.GetByID(ctx, id)
}
func (b *blogUsecase) Update(ctx context.Context, id primitive.ObjectID, blog *domain.BlogUpdateRequest) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.blogRepository.Update(ctx, id, blog)
}

func (b *blogUsecase) Delete(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, b.contextTimeout)
	defer cancel()
	return b.blogRepository.Delete(ctx, id)
}
