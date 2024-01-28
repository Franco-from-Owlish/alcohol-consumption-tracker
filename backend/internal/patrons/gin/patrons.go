package gin

import (
	"alcohol-consumption-tracker/internal/patrons"
	"alcohol-consumption-tracker/pkg/httpErrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		c.AbortWithStatusJSON(500, dbErr.JSON())
		return
	}
}
