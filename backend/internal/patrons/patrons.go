package patrons

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"alcohol-consumption-tracker/pkg/database"
	"alcohol-consumption-tracker/pkg/httpErrors"
)

type Patron struct {
	database.Model
	FirstName string `gorm:"size:48" json:"firstName"`
	LastName  string `gorm:"size:48" json:"lastName"`

	// Relationships
	Cocktails []cocktail.Cocktail `gorm:"many2many:patron_cocktails;" json:"cocktails,omitempty"` // has many

	// Optimisations
	TotalAlcohol float32 `json:"totalAlcohol"` // alcoholic content in ml
}

type PatronService interface {
	GetAllPatrons() (*[]Patron, error)
	CreatePatron(patron *Patron) error
	UpdatePatron(patron *Patron) error
	DeletePatron(patron *Patron) error
	DeletePatronByID(id string) error
	UpdateConsumption(patron *Patron, drink *cocktail.Cocktail) error
	GetPatronByID(id string) (*Patron, error)
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
