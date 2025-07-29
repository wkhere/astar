package geo

// Based on http://www.movable-type.co.uk/scripts/latlong.html

import "math"

const (
	piBy180     = math.Pi / 180
	earthRadius = 6371 //km
)

func H(p1, p2 Pt) float64 {
	f1 := toRadians(p1.Lat)
	f2 := toRadians(p2.Lat)
	df := toRadians(p2.Lat - p1.Lat)
	dl := toRadians(p2.Lng - p1.Lng)

	sp := math.Sin(df / 2)
	sl := math.Sin(dl / 2)
	a := sp*sp + math.Cos(f1)*math.Cos(f2)*sl*sl
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return math.Floor(earthRadius * c)
}

func toRadians(angle float64) float64 {
	return angle * piBy180
}
