package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	router.Use(cors.New(corsConfig))
	_ = router.SetTrustedProxies([]string{
		"localhost",
	})

	return router
}
