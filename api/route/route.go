package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(timeout time.Duration, db *mongo.Database, gin *gin.Engine){

	// publicRouter := gin.Group("")

	// NewSignupRouter(env, timeout, db, publicRouter)
	// NewLoginRouter(env, timeout, db, publicRouter)
	// NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	NewBlogRouter(timeout, db, protectedRouter)

}