package usecase

import (
	"Blog/domain"
	"Blog/internal/tokenutil"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	refreshTokenRepository domain.RefreshTokenRepository
	contextTimeout time.Duration

}
func NewRefreshTokenUsecase(userRepository domain.UserRepository, refreshTokenRepository domain.RefreshTokenRepository,timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		refreshTokenRepository: refreshTokenRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, id)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, userId string, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, userId, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}

func (rtu *refreshTokenUsecase) StoreRefreshToken(ctx context.Context, token *domain.RefreshToken) error {
	return rtu.refreshTokenRepository.StoreRefreshToken(ctx, token)
}

func (rtu *refreshTokenUsecase) GetStoredRefreshToken(ctx context.Context, userID primitive.ObjectID) (*domain.RefreshToken, error) {
	return rtu.refreshTokenRepository.GetStoredRefreshToken(ctx, userID)
}