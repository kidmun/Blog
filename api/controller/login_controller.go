package controller

import (
	"Blog/bootstrap"
	"Blog/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	RefreshTokenUseCase domain.RefreshTokenUsecase
	Env          *bootstrap.Env
}


func (lc *LoginController)Login(ctx *gin.Context){
	var request domain.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := lc.LoginUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		if err.Error() == "mongo: no documents in result"{
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}
	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	// encryptedRefreshToken, err := bcrypt.GenerateFromPassword(
	// 	[]byte(refreshToken),
	// 	bcrypt.DefaultCost,
	// )
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }
	expiryTime := time.Now().Add(time.Duration(lc.Env.RefreshTokenExpiryHour) * time.Hour)
	err = lc.RefreshTokenUseCase.StoreRefreshToken(ctx, &domain.RefreshToken{UserID: user.ID, ExpiresAt: expiryTime, Token: refreshToken})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to store refresh token"})
		return
	}
	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	ctx.JSON(http.StatusOK, loginResponse)
}