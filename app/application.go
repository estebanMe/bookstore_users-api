package app

import (
	"github.com/estebanMe/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

//StartApplication function starts the app
func StartApplication(){   
  mapUrls() 
  logger.Info("about to start the application...")
  router.Run(":8080")
}