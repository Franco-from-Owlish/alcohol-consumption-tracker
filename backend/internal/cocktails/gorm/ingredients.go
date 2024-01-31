package gorm

import (
	cocktaildb "alcohol-consumption-tracker/internal/cocktaildb/redis"
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ cocktail.IngredientsService = (*IngredientsService)(nil)

type IngredientsService struct {
	DB         *gorm.DB
	CocktailDB cocktaildb.CocktailDbService
}

func NewIngredientsService(db *gorm.DB, datastore *redis.Client) *IngredientsService {
	cocktailDbService := cocktaildb.NewCocktailDbService(datastore)

	return &IngredientsService{
		DB:         db,
		CocktailDB: *cocktailDbService,
	}
}

func (i IngredientsService) GetByName(name string) (*cocktail.Ingredient, error) {
	ing := cocktail.Ingredient{}
	errDB := i.DB.Where("name = ?", name).First(&ing).Error
	if errors.Is(errDB, gorm.ErrRecordNotFound) {
		ingData, err := i.CocktailDB.GetIngredient(name)
		if err != nil {
			return nil, err
		}
		ing = cocktail.Ingredient{
			Name: ingData.Name,
			Abv:  ingData.Abv,
		}
		if errSave := i.DB.Save(&ing).Error; errSave != nil {
			return nil, errSave
		}
	}
	return &ing, errDB
}
