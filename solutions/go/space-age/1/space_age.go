package space

type Planet string

const EarthYearSeconds float64 = 31557600

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / EarthYearSeconds
	case "Mercury":
		return seconds / (EarthYearSeconds * 0.2408467)
	case "Venus":
		return seconds / (EarthYearSeconds * 0.61519726)
	case "Mars":
		return seconds / (EarthYearSeconds * 1.8808158)
	case "Jupiter":
		return seconds / (EarthYearSeconds * 11.862615)
	case "Saturn":
		return seconds / (EarthYearSeconds * 29.447498)
	case "Uranus":
		return seconds / (EarthYearSeconds * 84.016846)
	case "Neptune":
		return seconds / (EarthYearSeconds * 164.79132)
	default:
		return 0
	}
}
