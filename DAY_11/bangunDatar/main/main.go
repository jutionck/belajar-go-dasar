package main

import (
	"bangunDatar/service"
	"bangunDatar/shape"
	"fmt"
)

func main() {
	segitiga := shape.Segitiga{2, 4}
	persegi := shape.Persegi{9}
	persegiPanjang := shape.PersegiPanjang{10, 5}
	trapesium := shape.Trapesium{2, 9, 5}

	serviceHitungLuas := service.NewHitungLuas()
	serviceHitungLuas.AddShapes(segitiga, persegi, persegiPanjang)
	serviceHitungLuas.AddAShape(trapesium)
	r := *(serviceHitungLuas.GetLuasTotal())
	fmt.Println(r)

	serviceHitungLuas.CetakLuasTotal()
}
