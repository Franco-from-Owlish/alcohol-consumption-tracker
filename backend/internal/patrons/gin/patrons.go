package gin

import (
	"alcohol-consumption-tracker/internal/patrons"
	"alcohol-consumption-tracker/pkg/httpErrors"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

const PatronNotFound = "Patron not found"
const PatronCollectionFailed = "Patron could not be collected"

func (s *Server) GetAll(c *gin.Context) {
	allPatrons, err := s.PatronService.GetAllPatrons()
	if err != nil {
		dbErr := httpErrors.NewDatabaseError(
			"Collecting all patrons failed", err)
		c.AbortWithStatusJSON(500, dbErr.JSON())
		return
	}
	c.JSON(http.StatusCreated, allPatrons)
}

func (s *Server) GetPatron(c *gin.Context) {
	id := c.Param("id")
	patron, err := s.PatronService.GetPatronByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pErr := httpErrors.NewParseError(PatronNotFound, err)
		c.AbortWithStatusJSON(pErr.HTTPStatus, pErr.JSON())
		return
	} else if err != nil {
		dbErr := httpErrors.NewDatabaseError(PatronCollectionFailed, err)
		c.AbortWithStatusJSON(dbErr.HTTPStatus, dbErr.JSON())
		return
	}
	c.JSON(http.StatusCreated, patron)
}

func (s *Server) RemovePatron(c *gin.Context) {
	id := c.Param("id")
	err := s.PatronService.DeletePatronByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pErr := httpErrors.NewParseError(PatronNotFound, err)
		c.AbortWithStatusJSON(pErr.HTTPStatus, pErr.JSON())
		return
	} else if err != nil {
		dbErr := httpErrors.NewDatabaseError(PatronCollectionFailed, err)
		c.AbortWithStatusJSON(dbErr.HTTPStatus, dbErr.JSON())
		return
	}

	c.JSON(http.StatusOK, "Deleted")
}

func (s *Server) AddCocktail(c *gin.Context) {
	id := c.Param("id")
	patron, err := s.PatronService.GetPatronByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		pErr := httpErrors.NewParseError(PatronNotFound, err)
		c.AbortWithStatusJSON(pErr.HTTPStatus, pErr.JSON())
		return
	} else if err != nil {
		dbErr := httpErrors.NewDatabaseError(PatronCollectionFailed, err)
		c.AbortWithStatusJSON(dbErr.HTTPStatus, dbErr.JSON())
		return
	}
	drinkName := c.Query("cocktail")
	drink, errCt := s.CocktailService.GetByName(drinkName)
	if errors.Is(errCt, gorm.ErrRecordNotFound) {
		pErr := httpErrors.NewParseError("Cocktail not found", errCt)
		c.AbortWithStatusJSON(pErr.HTTPStatus, pErr.JSON())
		return
	} else if errCt != nil {
		dbErr := httpErrors.NewDatabaseError(
			"Cocktail could not be collected", errCt)
		c.AbortWithStatusJSON(dbErr.HTTPStatus, dbErr.JSON())
		return
	}

	errAdd := s.PatronService.UpdateConsumption(patron, drink)
	if errAdd != nil {
		dbErr := httpErrors.NewDatabaseError(
			"Cocktail could not be added", errCt)
		c.AbortWithStatusJSON(dbErr.HTTPStatus, dbErr.JSON())
		return
	}

	c.JSON(http.StatusCreated, patron)
}

// CreatePatronInput JSON binding for expected input to CreateUser
type CreatePatronInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

// CreatePatron Creates a user
// Returns access keys and refresh keys.
func (s *Server) CreatePatron(c *gin.Context) {
	data := CreatePatronInput{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	patron := patrons.Patron{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	// Validate user input
	ve := patron.Validate()
	if !ve.IsEmpty() {
		c.AbortWithStatusJSON(ve.HTTPStatus, ve.JSON())
		return
	}
	if err := s.PatronService.CreatePatron(&patron); err != nil {
		dbErr := httpErrors.NewDatabaseError(
			"Creating patron failed", err)
		c.AbortWithStatusJSON(dbErr.HTTPStatus, dbErr.JSON())
		return
	}
	c.JSON(http.StatusCreated, patron)
}
