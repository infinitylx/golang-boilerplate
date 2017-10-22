package server

import (
	"../controllers"
	"../db"
	"../middlewares"
	"github.com/gin-gonic/gin"
)

// NewRouter return router
func NewRouter(dbConnect db.Connection) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())
	router.Use(middlewares.Connect(&dbConnect))
	v1 := router.Group("api/v1")
	{

		userController := new(controllers.UserController)

		v1.GET("/", userController.GetAll)
		v1.GET("/create", userController.CreateUser)

	}

	return router
}
