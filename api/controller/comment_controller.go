package controller

import (
	"Blog/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	CommentUsecase domain.CommentUsecase
}

func (cc *CommentController) Create(ctx *gin.Context) {
	
	var comment *domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err := cc.CommentUsecase.Create(ctx,comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "created successfully"})
}

func (cc *CommentController) Update(ctx *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	var comment *domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.CommentUsecase.Update(ctx, commentId, comment)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Comment not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Comment updated successfully",
	})
}

func (cc *CommentController) GetAll(ctx *gin.Context){
	comments, err := cc.CommentUsecase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

func (cc *CommentController) Delete(ctx *gin.Context) {
	commentId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	var comment *domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = cc.CommentUsecase.Delete(ctx, commentId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Comment not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Comment Dleeted successfully",
	})
}


func (cc *CommentController) GetByBlogID(ctx *gin.Context) {
	blogID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	comments, err := cc.CommentUsecase.GetByBlogID(ctx, blogID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)

}

func (cc *CommentController) GetByUserID(ctx *gin.Context) {
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	userIdPrim, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}

	comments, err := cc.CommentUsecase.GetByUserID(ctx, userIdPrim)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comments)

}

