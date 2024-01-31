package main

import (
	cocktail "alcohol-consumption-tracker/internal/cocktails"
	"alcohol-consumption-tracker/internal/patrons"
	_ "ariga.io/atlas-go-sdk/recordriver"
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"io"
	"os"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&patrons.Patron{},
		&cocktail.Cocktail{},
		&cocktail.Recipe{},
		&cocktail.Ingredient{},
		&cocktail.RecipeIngredient{},
	)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	_, _ = io.WriteString(os.Stdout, stmts)
}
