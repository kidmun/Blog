package controller

import (
	"Blog/domain"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
}

func (bc *BlogController) Create(ctx *gin.Context) {
	var blogInput domain.CreateBlogRequest
	if err := ctx.ShouldBindJSON(&blogInput); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	blog := &domain.Blog{
		Title:   blogInput.Title,
		Content: blogInput.Content,
		Tags:    blogInput.Tags,
		Date:    time.Now(),
	}
	validate := validator.New()
	if err := validate.Struct(blog); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	res, err := bc.BlogUsecase.Create(ctx, blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": res})
}

func (bc *BlogController) GetAll(ctx *gin.Context) {
	blogs, err := bc.BlogUsecase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) GetByID(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}

	blog, err := bc.BlogUsecase.GetByID(ctx, id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Blog not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, blog)

}

func (bc *BlogController) Update(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}

	var blog domain.BlogUpdateRequest

	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if blog.Title == nil && blog.Content == nil && blog.Tags == nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "At least one of Title, Content, or Tags must be provided"})
		return
	}

	err = bc.BlogUsecase.Update(ctx, id, &blog)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Blog not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task updated successfully",
	})
}

func (bc *BlogController) Delete(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	err = bc.BlogUsecase.Delete(ctx, id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Blog not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task deleted successfully",
	})
}

func
