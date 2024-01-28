package gin

import (
	"alcohol-consumption-tracker/internal/patrons"
	"alcohol-consumption-tracker/internal/patrons/gorm"
	"github.com/gin-gonic/gin"
	"net/http"

	database "gorm.io/gorm"
)

type Server struct {
	Router        *gin.RouterGroup
	PatronService patrons.PatronService
}

func NewServer(router *gin.RouterGroup, db *database.DB) *Server {
	userService := gorm.NewPatronService(db)

	s := &Server{
		router,
		userService,
	}

	s.addPublicRoutes()

	return s
}

func (s *Server) addPublicRoutes() {
	public := s.Router.Group("/")
	public.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
}
