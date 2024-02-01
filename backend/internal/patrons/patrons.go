package patrons

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"alcohol-consumption-tracker/pkg/database"
	"alcohol-consumption-tracker/pkg/ebac"
	"alcohol-consumption-tracker/pkg/httpErrors"
	"alcohol-consumption-tracker/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

type Patron struct {
	database.Model
	FirstName string  `gorm:"size:48" json:"firstName"`
	LastName  string  `gorm:"size:48" json:"lastName"`
	Sex       rune    `gorm:"default='M'" json:"sex"`
	Weight    float64 `gorm:"notnull" json:"weight"`

	// Relationships
	Cocktails []cocktail.Cocktail `gorm:"many2many:patron_cocktails;" json:"cocktails,omitempty"` // has many

	// Optimisations
	EBAC float64 `gorm:"default=0" json:"ebac"` // blood alcohol content
}

type PatronCocktails struct {
	CreatedAt  time.Time `json:"createdAt"`
	PatronID   int       `gorm:"primaryKey" json:"-"`
	CocktailID int       `gorm:"primaryKey" json:"-"`
}

type PatronService interface {
	GetAllPatrons() (*[]Patron, error)
	CreatePatron(patron *Patron) error
	UpdatePatron(patron *Patron) error
	DeletePatron(patron *Patron) error
	DeletePatronByID(id string) error
	UpdateConsumption(patron *Patron, drink *cocktail.Cocktail) error
	UpdateEBAC(patron *Patron) error
	GetPatronByID(id string) (*Patron, error)
}

func (p *Patron) Validate() httpErrors.ValidationError {
	ve := httpErrors.NewValidationError()

	if len(p.FirstName) == 0 {
		ve.AddRequiredField("first name")
	}
	if len(p.LastName) == 0 {
		ve.AddRequiredField("last name")
	}

	sexes := []rune{ebac.MALE, ebac.FEMALE}
	if !utils.Contains(p.Sex, sexes) {
		ve.AddInvalidFieldWithContext(
			"sex",
			fmt.Sprintf("must be one of %v", sexes),
		)
	}

	if p.Weight <= 0 {
		ve.AddInvalidField("weight")
	}
	if p.Weight > 635 { // see https://en.wikipedia.org/wiki/Jon_Brower_Minnoch
		ve.AddInvalidField("weight")
	}

	return *ve
}

func (p *Patron) MarshalJSON() ([]byte, error) {
	type PatronJSON Patron
	pJSON := struct {
		PatronJSON
		Sex string `json:"sex"`
	}{
		Sex:        string(p.Sex),
		PatronJSON: PatronJSON(*p),
	}
	return json.Marshal(&pJSON)
}
