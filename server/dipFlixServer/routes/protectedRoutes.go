package routes

import (
	controller "github.com/khdip/dip-flix/server/dipFlixServer/controllers"
	"github.com/khdip/dip-flix/server/dipFlixServer/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/movie/add", controller.AddMovie())
}
