package usecase

import (
	"Blog/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type logoutUsecase struct {
	refreshTokenRepository domain.RefreshTokenRepository
	contextTimeout time.Duration

}
func NewLogoutUsecase(refreshTokenRepository domain.RefreshTokenRepository,timeout time.Duration) domain.LogoutUsecase{
	return &logoutUsecase{
		refreshTokenRepository: refreshTokenRepository,
		contextTimeout: timeout,
	}
}

func (lu *logoutUsecase)  DeleteTokensByUserID(ctx context.Context, userID primitive.ObjectID) error {
	return lu.refreshTokenRepository.DeleteTokensByUserID(ctx, userID)
}



