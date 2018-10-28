package space

const (
	// EarthDays number of days in Earth year
	EarthDays = 365.25
	// EarthDaySeconds number of seconds in Earth day
	EarthDaySeconds = 86400
	// EarthYearSeconds number of seconds in Earth year
	EarthYearSeconds = EarthDays * EarthDaySeconds
)

// Planet type
type Planet string

var op = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age calculates Earth age in a given planet
func Age(s float64, p Planet) float64 {
	return s / (op[p] * EarthDays * EarthDaySeconds)
}
