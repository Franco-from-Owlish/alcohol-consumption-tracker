package cocktail

import (
	"alcohol-consumption-tracker/pkg/database"
)

type Cocktail struct {
	database.Model
	Name         string  `gorm:"size:48;unique" json:"name"`
	Recipe       Recipe  `json:"recipe,omitempty"`
	TotalAlcohol float64 `gorm:"default:0" json:"totalAlcohol"` // alcoholic content in ml
	OnMenu       bool    `gorm:"default:false" json:"onMenu"`
}

type Recipe struct {
	database.Model
	CocktailID uint `json:"-"`

	// Relationships
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;" json:"ingredients"` // has many
}

type RecipeIngredient struct {
	RecipeID     int     `gorm:"primaryKey" json:"-"`
	IngredientID int     `gorm:"primaryKey" json:"-"`
	Amount       float64 `json:"amount"`
	Unit         string  `gorm:"size:12" json:"unit"`
}

type CocktailsService interface {
	GetAll() (*[]Cocktail, error)
	GetRandom() (*Cocktail, error)
	GetByName(name string) (*Cocktail, error)
	GetCocktailIngredients(cocktail *Cocktail) error
	GetCocktailRecipe(cocktail *Cocktail) ([]map[string]interface{}, error)
	AddToMenu(name string) error
	UpdateCocktailAlcoholContent(cocktail *Cocktail) error
}
