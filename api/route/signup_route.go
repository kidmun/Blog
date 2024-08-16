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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
	group.GET("/verify", sc.VerifyEmail)
}