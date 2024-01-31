package gin

import (
	"alcohol-consumption-tracker/pkg/httpErrors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) GetAllCocktails(c *gin.Context) {
	cocktails, err := s.CocktailService.GetAll()
	if err != nil {
		databaseError := httpErrors.NewDatabaseError(
			"Fetching all cocktails failed",
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	c.JSON(http.StatusOK, cocktails)
}

func (s *Server) GetRandomCocktail(c *gin.Context) {
	cocktail, err := s.CocktailService.GetRandom()
	if err != nil {
		databaseError := httpErrors.NewDatabaseError(
			"Fetching random cocktail failed",
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	errUpdate := s.CocktailService.UpdateCocktailAlcoholContent(cocktail)
	if errUpdate != nil {
		databaseError := httpErrors.NewDatabaseError(
			"Updating random cocktail failed",
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	c.JSON(http.StatusOK, cocktail)
}

func (s *Server) GetCocktail(c *gin.Context) {
	name := c.Param("name")
	cocktail, err := s.CocktailService.GetByName(name)
	if err != nil {
		databaseError := httpErrors.NewDatabaseError(
			fmt.Sprintf("Fetching cocktail '%s' failed", name),
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	c.JSON(http.StatusOK, cocktail)
}

func (s *Server) GetCocktailIngredients(c *gin.Context) {
	name := c.Param("name")
	cocktail, err := s.CocktailService.GetByName(name)
	if err != nil {
		databaseError := httpErrors.NewDatabaseError(
			fmt.Sprintf("Fetching recipe for cocktail '%s' failed", name),
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	errRcp := s.CocktailService.GetCocktailIngredients(cocktail)
	if errRcp != nil {
		databaseError := httpErrors.NewDatabaseError(
			fmt.Sprintf("Fetching cocktail '%s' ingredients failed", name),
			errRcp,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	c.JSON(http.StatusOK, cocktail)
}

func (s *Server) GetCocktailRecipe(c *gin.Context) {
	name := c.Param("name")
	cocktail, err := s.CocktailService.GetByName(name)
	if err != nil {
		databaseError := httpErrors.NewDatabaseError(
			fmt.Sprintf("Fetching recipe for cocktail '%s' failed", name),
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	recipe, errRcp := s.CocktailService.GetCocktailRecipe(cocktail)
	if errRcp != nil {
		databaseError := httpErrors.NewDatabaseError(
			fmt.Sprintf("Fetching cocktail '%s' recipe failed", name),
			errRcp,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	c.JSON(http.StatusOK, recipe)
}

func (s *Server) AddCocktailToMenu(c *gin.Context) {
	name := c.Param("name")
	err := s.CocktailService.AddToMenu(name)
	if err != nil {
		databaseError := httpErrors.NewDatabaseError(
			fmt.Sprintf("Adding cocktail '%s' to menu failed", name),
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, databaseError.JSON())
		return
	}
	c.JSON(http.StatusOK, "OK")
}
