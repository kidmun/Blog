package route

import (
	"Blog/api/controller"
	"Blog/domain"
	"Blog/repository"
	"Blog/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLikeDislikeRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup){
	ldr := repository.NewLikeDislikeRepository(db, domain.CollectionLikeDislike)
	ldc := &controller.LikeDislikeController{
		LikeDislikeUsecase: usecase.NewLikeDislikeUsecase(ldr, timeout),	
	}
	group.POST("/add_like/:id", ldc.AddLike)
	group.POST("/remove_like/:id", ldc.RemoveLike)
	group.POST("/add_dislike/:id", ldc.AddDislike)
	group.POST("/remove_dislike/:id", ldc.RemoveDisLike)
}