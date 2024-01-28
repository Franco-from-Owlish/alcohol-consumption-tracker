package cocktail

import "gorm.io/gorm"

type Cocktail struct {
	gorm.Model
	FirstName    string `gorm:"size:48" json:"firstName"`
	LastName     string `gorm:"size:48" json:"lastName"`
	Recipe       Recipe
	TotalAlcohol float32 // alcoholic content in ml
	OnMenu       bool
}

type Recipe struct {
	ID         uint `gorm:"primarykey"`
	CocktailID uint

	// Relationships
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;" json:"addresses,omitempty"` // has many
}
