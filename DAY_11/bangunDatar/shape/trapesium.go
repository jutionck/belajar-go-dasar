package shape

type Trapesium struct {
	Qr     float64
	Ps     float64
	Tinggi float64
}

func (tr Trapesium) Luas() float64 {
	return (tr.Qr + tr.Ps) + tr.Tinggi/2
}

//Luas Trapesium Siku-siku PQRS = (QR + PS) x t / 2
