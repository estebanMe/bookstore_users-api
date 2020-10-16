package users

import (
	"fmt"
	"net/http"
	"strconv"

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
	fmt.Println("The user has been created", user)
	c.JSON(http.StatusCreated, result)
}

//GetUser GetUser
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusCreated, user)
}

//UpdateUser UpdateUser
func UpdateUser(c *gin.Context) {

	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)

}
