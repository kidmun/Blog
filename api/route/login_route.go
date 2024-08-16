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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rr := repository.NewRefreshTokenRepository(db, domain.CollectionRefreshToken)
	lc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		RefreshTokenUseCase: usecase.NewRefreshTokenUsecase(ur, rr, timeout),
		Env:           env,
	}
	group.POST("/login", lc.Login)
}