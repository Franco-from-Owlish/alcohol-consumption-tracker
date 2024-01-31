package gin

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"alcohol-consumption-tracker/internal/cocktails/gorm"
	"github.com/gin-gonic/gin"
	datastore "github.com/redis/go-redis/v9"
	database "gorm.io/gorm"
	"net/http"
)

type Server struct {
	Router          *gin.RouterGroup
	CocktailService cocktail.CocktailsService
}

func NewServer(router *gin.RouterGroup, db *database.DB, ds *datastore.Client) *Server {
	cocktailService := gorm.NewCocktailsService(db, ds)

	s := &Server{
		router,
		cocktailService,
	}

	s.addPublicRoutes()

	return s
}

func (s *Server) addPublicRoutes() {
	public := s.Router.Group("/")
	public.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "OK")
	})

	public.GET("/", s.GetAllCocktails)
	public.GET("/random", s.GetRandomCocktail)
	public.GET("/:name", s.GetCocktail)
	public.GET("/:name/recipe", s.GetCocktailRecipe)
	public.GET("/:name/add", s.AddCocktailToMenu)
}
