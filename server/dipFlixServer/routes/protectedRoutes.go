package routes

import (
	controller "github.com/khdip/dip-flix/server/dipFlixServer/controllers"
	"github.com/khdip/dip-flix/server/dipFlixServer/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.POST("/movie/add", controller.AddMovie(client))
	router.GET("/recommended", controller.GetRecommendedMovies(client))
	router.PATCH("/adminreview/:imdb_id", controller.AdminReviewUpdate(client))
}
