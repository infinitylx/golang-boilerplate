package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

//UserController mvc c part
type UserController struct{}

var userModel = new(models.User)

//GetAll return all users
func (u UserController) GetAll(c *gin.Context) {
	var user = models.User{}
	var result, err = user.GetAll()
	// err := conn.Use("eulist", "user").Find(nil).All(&result)

	// if err != nil {
	// 	return nil, err
	// }

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": result})
	return
}

// CreateUser creates user
func (u UserController) CreateUser(c *gin.Context) {
	// conn := c.MustGet("db").(*db.Connection)
	// defer conn.Close()

	nu := models.User{}
	nu.UserName = "Vasy"
	nu.SaveUser()
	// // err := conn.Use("eulist", "user").Insert(nu)
	// if err != nil {
	// 	c.JSON(500, gin.H{"message": "Can't insert", "error": err})
	// 	c.Abort()
	// 	return
	// }
	c.Redirect(http.StatusFound, "/api/v1")

}
