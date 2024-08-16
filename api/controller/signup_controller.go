package controller

import (
	"Blog/bootstrap"
	"Blog/domain"
	"Blog/internal/emailutil"
	"Blog/internal/tokenutil"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(ctx *gin.Context) {
	var request domain.SignupRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	existingUser, err := sc.SignupUsecase.GetUserByEmail(ctx, request.Email)

	if existingUser != nil && err == nil {
		ctx.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}
	if existingUser == nil && err.Error() != "mongo: no documents in result" {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	request.Password = string(encryptedPassword)
	user := domain.User{
		UserName: request.UserName,
		Email:    request.Email,
		Password: request.Password,
		Date:     time.Now(),
	}
	userID, err := sc.SignupUsecase.Create(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, userID.Hex(), sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	
	err = emailutil.SendVerificationEmail(accessToken, &user, sc.Env)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Signup successful! Please check your email to verify your account.",
	})
}

func (sc *SignupController) VerifyEmail(ctx *gin.Context) {
	tokenString := ctx.Query("token") // Assume token is passed as a query parameter

    if tokenString == "" {
        ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Token is required"})
        return
    }

	_, err := tokenutil.IsAuthorized(tokenString, sc.Env.AccessTokenSecret)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid or expired token"})
        return
    }
	stringId, err := tokenutil.ExtractIDFromToken(tokenString, sc.Env.AccessTokenSecret)

	if err != nil {
        ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Couldn't extract the user id from the token"})
        return
    }

	Id, err := primitive.ObjectIDFromHex(stringId)
    if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Couldn't parse the given id"})
		return
    }
	
	err = sc.SignupUsecase.UpdateUserVerificationStatus(ctx, Id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to update verification status"})
        return
    }

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})

}
