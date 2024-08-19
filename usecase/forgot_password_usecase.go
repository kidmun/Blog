package usecase

import (
	"Blog/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type forgotPasswordUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}
func NewForgotPasswordUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.ForgotPassowrdUsecase {
	return &forgotPasswordUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

func (fu *forgotPasswordUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, fu.contextTimeout)
	defer cancel()
	return fu.userRepository.GetByEmail(ctx, email)
}

func  (fu *forgotPasswordUsecase)  UpdatePassword(ctx context.Context, userID primitive.ObjectID, newPassword domain.ResetPasswordRequest) error{
	ctx, cancel := context.WithTimeout(ctx, fu.contextTimeout)
	defer cancel()
	return fu.userRepository.UpdatePassword(ctx, userID, newPassword)
}


