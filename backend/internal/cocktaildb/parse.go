package cocktaildb

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStrMeasure(measurement string) IngredientMeasurement {
	parts := strings.Split(strings.Trim(measurement, " "), " ")
	var value = 0.0
	var unit string
	partsCount := len(parts)

	if partsCount == 1 {
		parsedValue, errParse := strconv.ParseFloat(parts[0], 32)
		if errParse == nil {
			return IngredientMeasurement{Value: parsedValue, Unit: "unit"}
		} else {
			return IngredientMeasurement{Value: 0, Unit: parts[0]}
		}
	}

	for i := 0; i < partsCount-1; i++ {
		part := parts[i]
		if strings.Contains(part, "/") {
			value += fractionStringToFloat(part)
		} else {
			parsedValue, errParse := strconv.ParseFloat(part, 32)
			if errParse != nil {
				fmt.Printf("parse failed: %v", errParse)
				break
			}
			value += parsedValue
		}
	}

	switch parts[partsCount-1] {
	case "oz":
		value = OzToML(value)
		unit = "ml"
	case "shot":
		value = ShotToML(value)
		unit = "ml"
	default:
		unit = parts[partsCount-1]
	}

	return IngredientMeasurement{
		Value: value,
		Unit:  unit,
	}
}

func fractionStringToFloat(part string) float64 {
	fractionParts := strings.Split(part, "/")
	numerator, errParseN := strconv.ParseFloat(fractionParts[0], 32)
	if errParseN != nil {
		fmt.Printf("parsing numerator failed: %v", errParseN)
	}
	denominator, errParseD := strconv.ParseFloat(fractionParts[0], 32)
	if errParseD != nil {
		fmt.Printf("parsing denominator failed: %v", errParseD)
	}
	return numerator / denominator
}
