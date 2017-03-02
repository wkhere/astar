package geo

// Based on http://www.movable-type.co.uk/scripts/latlong.html

import "math"

const (
	piBy180     = math.Pi / 180
	earthRadius = 6371 //km
)

func H(p1, p2 Pt) float64 {
	phi1 := toRadian(p1.Lat)
	phi2 := toRadian(p2.Lat)
	dphi := toRadian(p2.Lat - p1.Lat)
	dlam := toRadian(p2.Lon - p1.Lon)

	sp := math.Sin(dphi / 2)
	sl := math.Sin(dlam / 2)
	a := sp*sp + math.Cos(phi1)*math.Cos(phi2)*sl*sl
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return math.Floor(earthRadius * c)
}

func toRadian(angle float64) float64 {
	return angle * piBy180
}
