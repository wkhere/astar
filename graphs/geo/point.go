package geo

type Pt struct {
	Lat, Lon float64
}

func Point(lat, lon float64) Pt {
	return Pt{lat, lon}
}

func AbsDeg(deg, min, sec float64) float64 {
	return deg + min/60 + sec/3600
}
