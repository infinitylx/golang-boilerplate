package server

import (
	"../controllers"
	"../middlewares"
	"github.com/gin-gonic/gin"
)

// NewRouter return router
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{

		userController := new(controllers.UserController)

		v1.GET("/", userController.GetAll)
		v1.GET("/create", userController.CreateUser)

	}

	return router
}
