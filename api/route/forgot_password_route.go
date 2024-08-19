package route

import (
	"Blog/api/controller"
	"Blog/bootstrap"
	"Blog/domain"
	"Blog/repository"
	"Blog/usecase"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewForgotPasswordRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.ForgotPasswordController{
		ForgotPasswordUsecase: usecase.NewForgotPasswordUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/forgot-password", rtc.RequestPasswordReset)
	group.POST("/reset-password", rtc.ResetPassword)
}