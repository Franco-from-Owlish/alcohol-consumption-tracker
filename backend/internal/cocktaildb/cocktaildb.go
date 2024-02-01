package cocktaildb

import "encoding/json"

type IngredientMeasurement struct {
	Value float64
	Unit  string
}

type Cocktail struct {
	Id          string                           `json:"Id"`
	Name        string                           `json:"name"`
	Alcoholic   bool                             `json:"alcoholic"`
	Ingredients map[string]IngredientMeasurement `json:"ingredients"`
}

type Ingredient struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Abv  float64 `json:"abv"`
}

type Service interface {
	GetCocktail(name string) (*Cocktail, error)
	GetRandomCocktail() (*Cocktail, error)
	GetIngredient(name string) (*Ingredient, error)
}

func JSON[T interface{}](data *T) []byte {
	value, _ := json.Marshal(data)
	return value
}

func (c *Cocktail) JSON() []byte {
	return JSON(c)
}

func (i *Ingredient) JSON() []byte {
	return JSON(i)
}
