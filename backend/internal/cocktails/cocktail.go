package cocktail

import (
	"gorm.io/gorm"
)

type Cocktail struct {
	gorm.Model
	Name         string `gorm:"size:48;unique" json:"name"`
	Recipe       Recipe
	TotalAlcohol float32 `gorm:"default:0"` // alcoholic content in ml
	OnMenu       bool    `gorm:"default:false"`
}

type Recipe struct {
	gorm.Model
	CocktailID uint

	// Relationships
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;" json:"ingredients"` // has many
}

type RecipeIngredient struct {
	RecipeID     int     `gorm:"primaryKey"`
	IngredientID int     `gorm:"primaryKey"`
	Amount       float32 `json:"amount"`
	Unit         string  `gorm:"size:12" json:"unit"`
}

type CocktailsService interface {
	GetAll() (*[]Cocktail, error)
	GetRandom() (*Cocktail, error)
	GetByName(name string) (*Cocktail, error)
	GetCocktailRecipe(cocktail *Cocktail) error
	AddToMenu(name string) error
	UpdateCocktailAlcoholContent(cocktail *Cocktail) error
}
