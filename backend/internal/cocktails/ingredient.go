package cocktail

import (
	"alcohol-consumption-tracker/pkg/database"
)

type Ingredient struct {
	database.Model
	Name string  `gorm:"size:48" json:"name"`
	Abv  float32 `json:"abv"`
}

type IngredientsService interface {
	GetByName(name string) (*Ingredient, error)
}
