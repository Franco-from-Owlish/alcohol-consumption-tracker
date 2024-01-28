package client

type IngredientResponse struct {
	IdIngredient   string `json:"idIngredient;omitempty"`
	StrIngredient  string `json:"strIngredient;omitempty"`
	StrDescription string `json:"strDescription;omitempty"`
	StrType        string `json:"strType;omitempty"`
	StrAlcohol     string `json:"strAlcohol;omitempty"`
	StrABV         int    `json:"strABV;omitempty"`
}

func GetIngredient() *RandomCocktailResponse {
	var resp map[string][]RandomCocktailResponse
	if err := GetBase(
		"search.php",
		Parameters{"i": "cola"},
		&resp,
	); err != nil {
		return nil
	}
	return &(resp["ingredients"][0])
}
