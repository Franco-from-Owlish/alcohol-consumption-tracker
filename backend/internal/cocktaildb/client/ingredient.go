package client

import (
	"alcohol-consumption-tracker/internal/cocktaildb"
)

type IngredientResponse struct {
	IdIngredient   string `json:"idIngredient"`
	StrIngredient  string `json:"strIngredient"`
	StrDescription string `json:"strDescription"`
	StrType        string `json:"strType"`
	StrAlcohol     string `json:"strAlcohol"`
	StrABV         int    `json:"strABV,string"`
}

func GetIngredient(name string) (*IngredientResponse, error) {
	var data map[string][]IngredientResponse
	resp, err := GetBase(
		"search.php",
		Parameters{"i": name},
		data,
	)
	if err != nil {
		return nil, err
	}
	return &resp["ingredients"][0], nil
}

func (i *IngredientResponse) ToIngredient() cocktaildb.Ingredient {
	return cocktaildb.Ingredient{
		Id:   i.IdIngredient,
		Name: i.StrIngredient,
		Abv:  float32(i.StrABV),
	}
}
