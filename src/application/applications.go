package application

import (
	"bookstore_users-api/src/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Start the aplication ... ")
	router.Run(":3001")

}
