package cocktaildb

const mlPerOz = 29.57353193
const mlPerShot = 45

func OzToML(oz float32) float32 {
	return oz * mlPerOz
}

func ShotToML(oz float32) float32 {
	return oz * mlPerShot
}
