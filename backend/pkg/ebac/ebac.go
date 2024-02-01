package ebac

import "time"

const (
	MALE   = 'M'
	FEMALE = 'F'
)

// alcoholWeightPerMilliLitre alcohol conversion factor between mass and volume
const alcoholWeightPerMilliLitre = 0.79

// beta rate of alcohol elimination in g/L/hr
const beta = 0.15

// distributionVolumeMale Average volume of distribution for males
const distributionVolumeMale = 0.71

// distributionVolumeFemale Average volume of distribution for females
const distributionVolumeFemale = 0.58

// CalculateEBAC
// see: https://en.wikipedia.org/wiki/Blood_alcohol_content
func CalculateEBAC(volume, weight float64, sex rune, now, consumed time.Time) float64 {
	A := volume * alcoholWeightPerMilliLitre
	var Vd float64
	if sex == MALE {
		Vd = distributionVolumeMale * weight
	} else {
		Vd = distributionVolumeFemale * weight
	}
	T := now.Sub(consumed).Hours()

	return A/Vd - beta*T
}
