package server

import (
	cocktails "alcohol-consumption-tracker/internal/cocktails/gin"
	patrons "alcohol-consumption-tracker/internal/patrons/gin"
	"alcohol-consumption-tracker/pkg/database/postgres"
	datastore "alcohol-consumption-tracker/pkg/database/redis"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	DS     *redis.Client
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
	redisStore := datastore.NewDataStore(
		os.Getenv("REDIS_USER"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_NAME"),
	)
	server.DS = redisStore.Open()

	server.Server = NewRouter()

	api := server.Server.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	patronsGroup := api.Group("/patron")
	cocktailsGroup := api.Group("/cocktail/")

	_ = cocktails.NewServer(cocktailsGroup, server.DB, server.DS)
	_ = patrons.NewServer(patronsGroup, server.DB, server.DS)

	return server
}
