package controller

import (
	"Blog/bootstrap"
	"Blog/domain"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(ctx *gin.Context) {
	var request domain.RefreshTokenRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id_string, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}
	id, err := primitive.ObjectIDFromHex(id_string)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Could not parse the given id."})
		return
	}
	user, err := rtc.RefreshTokenUsecase.GetUserByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "refresh token not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	storedToken, err := rtc.RefreshTokenUsecase.GetStoredRefreshToken(ctx, user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Refresh token not found"})
		return
	}
	if storedToken.Revoked {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Refresh token has been revoked"})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(user, user.ID.Hex(), rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	expiryTime := time.Now().Add(time.Duration(rtc.Env.RefreshTokenExpiryHour) * time.Hour)
	err = rtc.RefreshTokenUsecase.StoreRefreshToken(ctx, &domain.RefreshToken{UserID: user.ID, ExpiresAt: expiryTime, Token: refreshToken})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to store refresh token"})
		return
	}
	ctx.JSON(http.StatusOK, refreshTokenResponse)
}