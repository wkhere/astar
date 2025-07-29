package geo

type Pt struct {
	Lat, Lng float64
}

func Point(lat, lng float64) Pt {
	return Pt{lat, lng}
}

func AbsDeg(deg, min, sec float64) float64 {
	return deg + min/60 + sec/3600
}
