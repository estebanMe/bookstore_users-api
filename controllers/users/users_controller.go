package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser CreateUser
func CreateUser(c *gin.Context){
   c.String(http.StatusNotImplemented, "implement me!")
}

//GetUser GetUser
func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "implement me!")
}
