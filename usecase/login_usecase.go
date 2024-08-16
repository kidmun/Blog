package usecase

import (
	"Blog/domain"
	"Blog/internal/tokenutil"
	"context"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.LoginUsecase {
return &loginUsecase{
	userRepository: userRepository,
	contextTimeout: contextTimeout,
}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, user.ID.Hex(), secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
