package cocktaildb

type Cocktail struct {
	Id          string
	Name        string
	Alcoholic   bool
	Ingredients map[string]float32
}

type Ingredient struct {
	Id   string
	Name string
	Abv  int
}

type Service interface {
	GetRandomCocktail() (Cocktail, error)
	GetIngredient() (Ingredient, error)
}
