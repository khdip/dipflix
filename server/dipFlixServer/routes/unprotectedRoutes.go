package routes

import (
	controller "github.com/khdip/dip-flix/server/dipFlixServer/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUnprotectedRoutes(router *gin.Engine) {
	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	router.PATCH("/adminreview/:imdb_id", controller.AdminReviewUpdate())
}
