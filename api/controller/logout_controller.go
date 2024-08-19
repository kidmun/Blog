package controller

import (
	"Blog/bootstrap"
	"Blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogoutController struct {
	LogoutUsecase domain.LogoutUsecase
	Env           *bootstrap.Env
}


func (lc *LogoutController) Logout(ctx *gin.Context) {
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	id, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	err = lc.LogoutUsecase.DeleteTokensByUserID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Logout successful",
	})

}