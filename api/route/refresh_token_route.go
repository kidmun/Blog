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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewRefreshTokenRepository(db, domain.CollectionRefreshToken)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, rr, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}