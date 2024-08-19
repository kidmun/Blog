package controller

import (
	"Blog/bootstrap"
	"Blog/domain"
	"Blog/internal/emailutil"
	"Blog/internal/tokenutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type ForgotPasswordController struct {
	ForgotPasswordUsecase        domain.ForgotPassowrdUsecase
	Env                 *bootstrap.Env
}

func (fpc *ForgotPasswordController) RequestPasswordReset(ctx *gin.Context) {
	var request domain.PasswordResetRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := fpc.ForgotPasswordUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	resetToken, err := tokenutil.GeneratePasswordResetToken(user, fpc.Env.PasswordResetTokenSecret, fpc.Env.PasswordResetTokenExpiryHour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = emailutil.SendPasswordResetEmail(resetToken, user, fpc.Env)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}

func (fpc *ForgotPasswordController) ResetPassword(ctx *gin.Context) {
	userID, exists := ctx.Get("x-user-id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User ID not found in context"})
		return
	}
	tokenString := ctx.Query("token") // Assume token is passed as a query parameter
    if tokenString == "" {
        ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Token is required"})
        return
    }
	
	id, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	var request domain.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	_, err = tokenutil.IsAuthorized(tokenString, fpc.Env.PasswordResetTokenSecret)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid or expired token"})
        return
    }
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	request.NewPassword = string(encryptedPassword)
	err = fpc.ForgotPasswordUsecase.UpdatePassword(ctx, id, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password successfully reset"})

}