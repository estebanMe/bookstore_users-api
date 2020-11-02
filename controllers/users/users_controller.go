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

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userID, nil
}

//Create create user controller
func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: handle user creation error
		return
	}
	fmt.Println("The user has been created", user)
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//Get Get user controller
func Get(c *gin.Context) {
	userID, idErr := getUserID(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusCreated, user.Marshall(c.GetHeader("X-Public") == "true"))
}

//Update Update user controller
func Update(c *gin.Context) {

	userID, idErr := getUserID(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
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

	result, err := services.UsersService.UpdateUser(isPartial, user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

//Delete delete user controller
func Delete(c *gin.Context) {

	userID, idErr := getUserID(c.Param("user_id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
 
	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "delete"})

}

//Search serch from query status
func Search(c *gin.Context ){
  status := c.Query("status")
 
  users, err := services.UsersService.SearchUser(status)

  if err != nil {
	  c.JSON(err.Status, err)
	  return
  }
  

  c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}