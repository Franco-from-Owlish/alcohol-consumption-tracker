package cocktaildb

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStrMeasure(measurement string) IngredientMeasurement {
	parts := strings.Split(strings.Trim(measurement, " "), " ")
	var value float32 = 0.0
	var unit string
	partsCount := len(parts)

	if partsCount == 1 {
		parsedValue, errParse := strconv.ParseFloat(parts[0], 32)
		if errParse == nil {
			return IngredientMeasurement{Value: float32(parsedValue), Unit: "unit"}
		} else {
			return IngredientMeasurement{Value: 0, Unit: parts[0]}
		}
	}

	for i := 0; i < partsCount-1; i++ {
		switch parts[i] {
		case "1/2":
			value += 0.5
		case "1/3":
			value += 1 / 3
		case "1/4":
			value += 0.25
		case "1/6":
			value += 1 / 6
		case "1/8":
			value += 0.125
		default:
			parsedValue, errParse := strconv.ParseFloat(parts[i], 32)
			if errParse != nil {
				fmt.Printf("parse failed: %v", errParse)
				break
			}
			value += float32(parsedValue)
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
