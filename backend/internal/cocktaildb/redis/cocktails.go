package redis

import (
	"alcohol-consumption-tracker/internal/cocktaildb"
	"alcohol-consumption-tracker/internal/cocktaildb/client"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var _ cocktaildb.Service = (*CocktailDbService)(nil)

var cacheTime = time.Duration(2 * 3.6e12) // 2 hours

type CocktailDbService struct {
	DS *redis.Client
}

func NewCocktailDbService(datastore *redis.Client) *CocktailDbService {
	return &CocktailDbService{DS: datastore}
}

func cocktailKey(name string) string {
	return fmt.Sprintf("cocktail-%s", name)
}

func ingredientKey(name string) string {
	return fmt.Sprintf("ingredient-%s", name)
}

func (c CocktailDbService) GetCocktail(name string) (*cocktaildb.Cocktail, error) {
	ctx := context.Background()
	data, errRtv := c.DS.Get(ctx, cocktailKey(name)).Result()
	if errors.Is(errRtv, redis.Nil) {
		value := client.GetCocktail(name).ToCocktail()
		errSet := c.DS.Set(ctx, cocktailKey(name), value.JSON(), cacheTime).Err()
		if errSet != nil {
			return nil, errSet
		}
		return &value, nil
	}
	var value cocktaildb.Cocktail
	err := json.Unmarshal([]byte(data), &value)
	return &value, err
}

func (c CocktailDbService) GetRandomCocktail() (*cocktaildb.Cocktail, error) {
	ctx := context.Background()
	var exists = true
	var value cocktaildb.Cocktail
	for exists == true {
		value = client.GetRandomCocktail().ToCocktail()
		result, err := c.DS.Exists(ctx, cocktailKey(value.Name)).Result()
		if err != nil {
			return nil, err
		}
		exists = result >= 1
	}
	errSet := c.DS.Set(ctx, cocktailKey(value.Name), value.JSON(), cacheTime).Err()
	if errSet != nil {
		return nil, errSet
	}
	return &value, nil
}

func (c CocktailDbService) GetIngredient(name string) (*cocktaildb.Ingredient, error) {
	ctx := context.Background()
	data, errRtv := c.DS.Get(ctx, ingredientKey(name)).Result()
	if errors.Is(errRtv, redis.Nil) {
		response, errFetch := client.GetIngredient(name)
		if errFetch != nil {
			return nil, errFetch
		}
		value := response.ToIngredient()
		errSet := c.DS.Set(ctx, ingredientKey(name), value.JSON(), cacheTime).Err()
		if errSet != nil {
			return nil, errSet
		}
		return &value, nil
	} else if errRtv != nil {
		return nil, errRtv
	}
	var value cocktaildb.Ingredient
	err := json.Unmarshal([]byte(data), &value)
	fmt.Printf("\nIngredient:\n%v\n", string(value.JSON()))
	return &value, err
}
