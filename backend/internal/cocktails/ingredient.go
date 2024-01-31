package cocktail

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name string  `gorm:"size:48" json:"name"`
	Abv  float32 `json:"abv"` // save percentage as int * 100, 40{%} = 40000
}

type IngredientsService interface {
	GetByName(name string) (*Ingredient, error)
}

func (i *Ingredient) AbvAsDecimal() float32 {
	return float32(i.Abv / 10000)
}
