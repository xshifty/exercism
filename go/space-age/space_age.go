package space

type Planet string

var lookup = map[Planet]float64{
	"Earth":   31557600.0,
	"Mercury": 7600543.81992,
	"Venus":   19414149.0522,
	"Mars":    59354032.6901,
	"Jupiter": 374355659.124,
	"Saturn":  929292362.885,
	"Uranus":  2651370019.33,
	"Neptune": 5200418560.03}

func Age(seconds float64, planet Planet) float64 {
	proportion, ok := lookup[planet]

	if !ok {
		return 0.
	}

	return seconds / proportion
}
