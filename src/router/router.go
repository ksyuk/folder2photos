package router

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.GET("/access-token", Oauth2)

	router.GET("/callback", Oauth2Callback)

	router.Run(":8080")
}
