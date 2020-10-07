package users

import (
	"fmt"
	"net/http"

	"github.com/estebanMe/bookstore_users-api/domain/users"
	"github.com/estebanMe/bookstore_users-api/services"
	"github.com/estebanMe/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//CreateUser CreateUser
func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle user creation error
		return
	}
	fmt.Println("user created", user)
	c.JSON(http.StatusCreated, result)
}

//GetUser GetUser
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
