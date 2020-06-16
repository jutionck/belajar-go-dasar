package shape

type Segitiga struct {
	Alas   float64
	Tinggi float64
}

func (s Segitiga) Luas() float64 {
	return s.Alas * s.Tinggi / 2
}
