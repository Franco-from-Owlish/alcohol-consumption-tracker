package server

import (
	patrons "alcohol-consumption-tracker/internal/patrons/gin"
	"alcohol-consumption-tracker/pkg/database/postgres"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Server *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	postgresDatabase := postgres.NewDatabase(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	var dbErr error
	server.DB, dbErr = postgresDatabase.Open()
	if dbErr != nil {
		log.Fatal("Connection to db failed for server")
	}

	server.Server = NewRouter()

	server.Server.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	api := server.Server.Group("/api")
	patronsGroup := api.Group("/patron")
	//cocktailsGroup := api.Group("/cocktails/")

	_ = patrons.NewServer(patronsGroup, server.DB)
	//_ = cocktails.NewServer(cocktailsGroup, postgresDatabase)

	return server
}
