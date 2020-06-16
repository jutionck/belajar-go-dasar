package service

import (
	"bangunDatar/shape"
	"fmt"
)

type perhitunganLuasBangunDatar struct {
	kumpulanBangunDatar []shape.HitungBangunDatar
	hasil               *float64
}

//Constructor
func NewHitungLuas() *perhitunganLuasBangunDatar {
	hasilPtr := new(float64)
	kumpulan := make([]shape.HitungBangunDatar, 0)
	return &perhitunganLuasBangunDatar{kumpulanBangunDatar: kumpulan, hasil: hasilPtr}
}

func (p *perhitunganLuasBangunDatar) AddAShape(shape shape.HitungBangunDatar) {
	p.kumpulanBangunDatar = append(p.kumpulanBangunDatar, shape)
}

func (p *perhitunganLuasBangunDatar) AddShapes(shapes ...shape.HitungBangunDatar) {
	for _, s := range shapes {
		p.AddAShape(s)
	}
}

func (p *perhitunganLuasBangunDatar) luasBangunDatar(bd shape.HitungBangunDatar) float64 {
	hasil := bd.Luas()
	return hasil
}

//Instant Print
func (p *perhitunganLuasBangunDatar) CetakLuasTotal() {
	p.GetLuasTotal()
	p.cetak()
}

//For Further Calculation, provide the result
func (p *perhitunganLuasBangunDatar) GetLuasTotal() *float64 {
	var tot float64
	for _, value := range p.kumpulanBangunDatar {
		tot += p.luasBangunDatar(value)
	}
	*(p.hasil) = tot
	return p.hasil
}

func (p *perhitunganLuasBangunDatar) cetak() {
	fmt.Println(*p.hasil)
}
