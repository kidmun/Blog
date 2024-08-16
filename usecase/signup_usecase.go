package usecase

import (
	"Blog/domain"
	"Blog/internal/tokenutil"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}
func (su *signupUsecase) Create(c context.Context, user *domain.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, userId string, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, userId, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
func (su *signupUsecase) UpdateUserVerificationStatus(ctx context.Context, userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.userRepository.UpdateUserVerificationStatus(ctx, userID)
}
