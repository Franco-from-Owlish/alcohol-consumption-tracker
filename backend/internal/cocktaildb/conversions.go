package cocktaildb

const mlPerOz = 29.57353193
const mlPerShot = 45

func OzToML(oz float64) float64 {
	return oz * mlPerOz
}

func ShotToML(oz float64) float64 {
	return oz * mlPerShot
}
