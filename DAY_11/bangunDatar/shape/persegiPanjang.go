package shape

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

func (pj PersegiPanjang) Luas() float64 {
	return pj.Panjang * pj.Lebar
}
