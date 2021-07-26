package Controllers

import (
	"fmt"
	"github.com/exercise2/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var person []Models.Person
	err := Models.GetAllUsers(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var person Models.Person
	c.BindJSON(&person)
	err := Models.CreateUser(&person)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Models.Person
	err := Models.GetUserByID(&person, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var person Models.Person
	id := c.Params.ByName("id")
	err := Models.GetUserByID(&person, id)
	if err != nil {
		c.JSON(http.StatusNotFound, person)
	}
	c.BindJSON(&person)
	err = Models.UpdateUser(&person, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var person Models.Person
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&person, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}