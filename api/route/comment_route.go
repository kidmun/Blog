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

func NewCommentRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup){
	cr := repository.NewCommentRepository(db, domain.CollectionComment)
	cc := &controller.CommentController{
		CommentUsecase: usecase.NewCommentUsecase(cr, timeout),
	}
	group.POST("/blog/comments", cc.Create)
	group.GET("/blog/comments", cc.GetAll)
	group.GET("/blog/comments/:id", cc.GetByBlogID)
	group.GET("/user/comments/:id", cc.GetByUserID)
	group.PUT("/blog/comments/:id", cc.Update)
	group.DELETE("/blog/comments/:id", cc.Delete)
}