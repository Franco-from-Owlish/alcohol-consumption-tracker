package cocktaildb

import "strconv"

const mlPerOz = 29.57353193

func OzToML(oz string) float32 {
	value, _ := strconv.ParseFloat(oz, 32)
	return float32(value * mlPerOz)
}
