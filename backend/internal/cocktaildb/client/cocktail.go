package client

import (
	"alcohol-consumption-tracker/internal/cocktaildb"
)

type CocktailResponse struct {
	IdDrink         string `json:"idDrink"`
	StrDrink        string `json:"strDrink"`
	StrAlcoholic    string `json:"strAlcoholic"`
	StrGlass        string `json:"strGlass"`
	StrInstructions string `json:"strInstructions"`
	StrDrinkThumb   string `json:"strDrinkThumb"`
	StrIngredient1  string `json:"strIngredient1"`
	StrIngredient2  string `json:"strIngredient2"`
	StrIngredient3  string `json:"strIngredient3"`
	StrIngredient4  string `json:"strIngredient4"`
	StrIngredient5  string `json:"strIngredient5"`
	StrIngredient6  string `json:"strIngredient6"`
	StrIngredient7  string `json:"strIngredient7"`
	StrIngredient8  string `json:"strIngredient8"`
	StrIngredient9  string `json:"strIngredient9"`
	StrIngredient10 string `json:"strIngredient10"`
	StrIngredient11 string `json:"strIngredient11"`
	StrIngredient12 string `json:"strIngredient12"`
	StrIngredient13 string `json:"strIngredient13"`
	StrIngredient14 string `json:"strIngredient14"`
	StrIngredient15 string `json:"strIngredient15"`
	StrMeasure1     string `json:"strMeasure1"`
	StrMeasure2     string `json:"strMeasure2"`
	StrMeasure3     string `json:"strMeasure3"`
	StrMeasure4     string `json:"strMeasure4"`
	StrMeasure5     string `json:"strMeasure5"`
	StrMeasure6     string `json:"strMeasure6"`
	StrMeasure7     string `json:"strMeasure7"`
	StrMeasure8     string `json:"strMeasure8"`
	StrMeasure9     string `json:"strMeasure9"`
	StrMeasure10    string `json:"strMeasure10"`
	StrMeasure11    string `json:"strMeasure11"`
	StrMeasure12    string `json:"strMeasure12"`
	StrMeasure13    string `json:"strMeasure13"`
	StrMeasure14    string `json:"strMeasure14"`
	StrMeasure15    string `json:"strMeasure15"`
}

func GetCocktail(name string) *CocktailResponse {
	var data map[string][]CocktailResponse
	resp, err := GetBase(
		"search.php",
		Parameters{"s": name},
		data,
	)
	if err != nil {
		return nil
	}
	return &resp["drinks"][0]
}

func GetRandomCocktail() *CocktailResponse {
	var data map[string][]CocktailResponse
	resp, err := GetBase(
		"random.php", Parameters{}, data)
	if err != nil {
		return nil
	}
	return &resp["drinks"][0]
}

func (r *CocktailResponse) ToCocktail() cocktaildb.Cocktail {
	cocktail := cocktaildb.Cocktail{
		Id:          r.IdDrink,
		Name:        r.StrDrink,
		Alcoholic:   r.StrAlcoholic == "Alcoholic",
		Ingredients: map[string]cocktaildb.IngredientMeasurement{},
	}

	r.parseIngredients(&cocktail)

	return cocktail
}

func (r *CocktailResponse) parseIngredients(cocktail *cocktaildb.Cocktail) {
	if r.StrIngredient1 != "" {
		cocktail.Ingredients[r.StrIngredient1] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure1)
	}
	if r.StrIngredient2 != "" {
		cocktail.Ingredients[r.StrIngredient2] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure2)
	}
	if r.StrIngredient3 != "" {
		cocktail.Ingredients[r.StrIngredient3] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure3)
	}
	if r.StrIngredient4 != "" {
		cocktail.Ingredients[r.StrIngredient4] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure4)
	}
	if r.StrIngredient5 != "" {
		cocktail.Ingredients[r.StrIngredient5] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure5)
	}
	if r.StrIngredient6 != "" {
		cocktail.Ingredients[r.StrIngredient6] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure6)
	}
	if r.StrIngredient7 != "" {
		cocktail.Ingredients[r.StrIngredient7] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure7)
	}
	if r.StrIngredient8 != "" {
		cocktail.Ingredients[r.StrIngredient8] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure8)
	}
	if r.StrIngredient9 != "" {
		cocktail.Ingredients[r.StrIngredient9] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure9)
	}
	if r.StrIngredient10 != "" {
		cocktail.Ingredients[r.StrIngredient10] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure10)
	}
	if r.StrIngredient11 != "" {
		cocktail.Ingredients[r.StrIngredient11] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure11)
	}
	if r.StrIngredient12 != "" {
		cocktail.Ingredients[r.StrIngredient12] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure12)
	}
	if r.StrIngredient13 != "" {
		cocktail.Ingredients[r.StrIngredient13] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure13)
	}
	if r.StrIngredient14 != "" {
		cocktail.Ingredients[r.StrIngredient14] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure14)
	}
	if r.StrIngredient15 != "" {
		cocktail.Ingredients[r.StrIngredient15] =
			cocktaildb.ConvertStrMeasure(r.StrMeasure15)
	}
}
