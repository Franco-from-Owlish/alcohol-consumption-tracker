package client

type RandomCocktailResponse struct {
	IdDrink         string `json:"idDrink;omitempty"`
	StrDrink        string `json:"strDrink;omitempty"`
	StrAlcoholic    string `json:"strAlcoholic;omitempty"`
	StrGlass        string `json:"strGlass;omitempty"`
	StrInstructions string `json:"strInstructions;omitempty"`
	StrDrinkThumb   string `json:"strDrinkThumb;omitempty"`
	StrIngredient1  string `json:"strIngredient1;omitempty"`
	StrIngredient2  string `json:"strIngredient2;omitempty"`
	StrIngredient3  string `json:"strIngredient3;omitempty"`
	StrIngredient4  string `json:"strIngredient4;omitempty"`
	StrIngredient5  string `json:"strIngredient5;omitempty"`
	StrIngredient6  string `json:"strIngredient6;omitempty"`
	StrIngredient7  string `json:"strIngredient7;omitempty"`
	StrIngredient8  string `json:"strIngredient8;omitempty"`
	StrIngredient9  string `json:"strIngredient9;omitempty"`
	StrIngredient10 string `json:"strIngredient10;omitempty"`
	StrIngredient11 string `json:"strIngredient11;omitempty"`
	StrIngredient12 string `json:"strIngredient12;omitempty"`
	StrIngredient13 string `json:"strIngredient13;omitempty"`
	StrIngredient14 string `json:"strIngredient14;omitempty"`
	StrIngredient15 string `json:"strIngredient15;omitempty"`
	StrMeasure1     string `json:"strMeasure1;omitempty"`
	StrMeasure2     string `json:"strMeasure2;omitempty"`
	StrMeasure3     string `json:"strMeasure3;omitempty"`
	StrMeasure4     string `json:"strMeasure4;omitempty"`
	StrMeasure5     string `json:"strMeasure5;omitempty"`
	StrMeasure6     string `json:"strMeasure6;omitempty"`
	StrMeasure7     string `json:"strMeasure7;omitempty"`
	StrMeasure8     string `json:"strMeasure8;omitempty"`
	StrMeasure9     string `json:"strMeasure9;omitempty"`
	StrMeasure10    string `json:"strMeasure10;omitempty"`
	StrMeasure11    string `json:"strMeasure11;omitempty"`
	StrMeasure12    string `json:"strMeasure12;omitempty"`
	StrMeasure13    string `json:"strMeasure13;omitempty"`
	StrMeasure14    string `json:"strMeasure14;omitempty"`
	StrMeasure15    string `json:"strMeasure15;omitempty"`
}

func GetRandomCocktail() *RandomCocktailResponse {
	var resp map[string][]RandomCocktailResponse
	if err := GetBase("random.php", Parameters{}, &resp); err != nil {
		return nil
	}
	return &(resp["drinks"][0])
}
