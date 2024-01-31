package gin

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails/gorm"
	"alcohol-consumption-tracker/internal/patrons"
	"alcohol-consumption-tracker/internal/patrons/gorm"
	"github.com/gin-gonic/gin"
	"net/http"

	datastore "github.com/redis/go-redis/v9"
	database "gorm.io/gorm"
)

type Server struct {
	Router          *gin.RouterGroup
	PatronService   patrons.PatronService
	CocktailService cocktail.CocktailsService
}

func NewServer(router *gin.RouterGroup, db *database.DB, ds *datastore.Client) *Server {
	patronService := gorm.NewPatronService(db)
	cocktailService := cocktail.NewCocktailsService(db, ds)

	s := &Server{
		router,
		patronService,
		*cocktailService,
	}

	s.addPublicRoutes()

	return s
}

func (s *Server) addPublicRoutes() {
	public := s.Router.Group("/")
	public.GET("/health", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	public.GET("/", s.GetAll)
	public.POST("/", s.CreatePatron)

	public.GET("/:id", s.GetPatron)
	public.DELETE("/:id", s.RemovePatron)
	public.PUT("/:id", s.AddCocktail)
}
