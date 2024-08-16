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

func NewBlogRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup){
	br := repository.NewBlogRepository(db, domain.CollectionBlog)
	bc := &controller.BlogController{
		BlogUsecase: usecase.NewBlogUsecase(br, timeout),
	}
	group.POST("/blogs", bc.Create)
	group.GET("/blogs", bc.GetAll)
	group.GET("/blogs/:id", bc.GetByID)
	group.PUT("/blogs/:id", bc.Update)
	group.DELETE("/blogs/:id", bc.Delete)
}