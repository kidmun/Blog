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

func NewLogoutRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {

	rr := repository.NewRefreshTokenRepository(db, domain.CollectionRefreshToken)
	lc := controller.LogoutController{
		LogoutUsecase: usecase.NewLogoutUsecase(rr, timeout),
		Env:           env,
	}
	group.POST("/logout", lc.Logout)
}