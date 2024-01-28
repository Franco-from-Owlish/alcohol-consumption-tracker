package patrons

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"alcohol-consumption-tracker/pkg/httpErrors"
	"gorm.io/gorm"
)

type Patron struct {
	gorm.Model
	FirstName string `gorm:"size:48" json:"firstName"`
	LastName  string `gorm:"size:48" json:"lastName"`

	// Relationships
	Cocktails []cocktail.Cocktail `gorm:"many2many:patron_cocktails;" json:"addresses,omitempty"` // has many

	// Optimisations
	TotalAlcohol float32 // alcoholic content in ml
}

type PatronService interface {
	CreatePatron(patron *Patron) error
	UpdatePatron(patron *Patron) error
	GetPatronByID(id string) (*Patron, error)
	UpdateConsumption(id uint) (*Patron, error)
}

func (p *Patron) Validate() httpErrors.ValidationError {
	ve := httpErrors.NewValidationError()

	if len(p.FirstName) == 0 {
		ve.AddRequiredField("first name")
	}
	if len(p.LastName) == 0 {
		ve.AddRequiredField("first name")
	}

	return *ve
}
