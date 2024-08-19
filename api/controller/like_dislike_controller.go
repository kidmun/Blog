package controller

import (
	"Blog/domain"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeDislikeController struct {
	LikeDislikeUsecase domain.LikeDislikeUsecase
}


func (ldc *LikeDislikeController) AddLike(ctx *gin.Context){
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	blogID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	id, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	var likeDislike  = domain.LikeDislike{
		UserID: id,
		BlogID: blogID,
		IsLike: true,
	}
	err = ldc.LikeDislikeUsecase.AddLike(ctx, &likeDislike)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Like added successfully",
	})
}

func (ldc *LikeDislikeController) RemoveLike(ctx *gin.Context){
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	blogID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	userId, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	err = ldc.LikeDislikeUsecase.RemoveLike(ctx, blogID, userId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Comment not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Like added successfully",
	})
}
func (ldc *LikeDislikeController) AddDislike(ctx *gin.Context){
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	blogID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	id, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	var likeDislike  = domain.LikeDislike{
		UserID: id,
		BlogID: blogID,
		IsLike: false,
	}
	err = ldc.LikeDislikeUsecase.AddLike(ctx, &likeDislike)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Like added successfully",
	})
}

func (ldc *LikeDislikeController) RemoveDisLike(ctx *gin.Context){
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	blogID, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	userId, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	err = ldc.LikeDislikeUsecase.RemoveLike(ctx, blogID, userId)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "Comment not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Like added successfully",
	})
}