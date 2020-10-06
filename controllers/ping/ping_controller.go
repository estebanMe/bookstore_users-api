package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping Verify server status with a ping
func Ping(c *gin.Context){
	c.String(http.StatusOK, "pong")
}