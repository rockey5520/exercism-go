package space

// Planet is the input type
type Planet string

const earthrevolution = 31557600

// Age calculates the age
func Age(age float64, planet Planet) float64 {

	Planets := map[Planet]float64{
		"Earth":   earthrevolution,
		"Mercury": earthrevolution * 0.2408467,
		"Venus":   earthrevolution * 0.61519726,
		"Mars":    earthrevolution * 1.8808158,
		"Jupiter": earthrevolution * 11.862615,
		"Saturn":  earthrevolution * 29.447498,
		"Uranus":  earthrevolution * 84.016846,
		"Neptune": earthrevolution * 164.79132,
	}

	return age / Planets[planet]
}
