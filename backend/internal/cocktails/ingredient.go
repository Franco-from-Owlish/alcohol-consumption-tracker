package cocktail

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name string `gorm:"size:48" json:"name"`
	Abv  int    `json:"abv"` // save percentage as int * 100, 40% = 40000
}

func (i *Ingredient) GetAbvAsDecimal() float32 {
	return float32(i.Abv / 10000)
}
