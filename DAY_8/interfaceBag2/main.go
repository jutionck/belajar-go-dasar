package main

import (
	"fmt"
	"math"
)

type bangunDatar interface {
	lusasBangunDatar() float64
	kelilingBangunDatar() float64
}

type lingkaran struct {
	jariJari float64
}

type persegi struct {
	sisi float64
}

func (l lingkaran) lusasBangunDatar() float64 {
	return math.Pi * math.Pow(l.jariJari, 2)
}

func (l lingkaran) kelilingBangunDatar() float64 {
	return math.Pi * l.jariJari * 2
}

func (p persegi) lusasBangunDatar() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) kelilingBangunDatar() float64 {
	return 4 * p.sisi
}

func printTwoDimentionalFigure(b bangunDatar) {
	area := fmt.Sprintf("%.2f", b.lusasBangunDatar())
	around := fmt.Sprintf("%.2f", b.kelilingBangunDatar())
	fmt.Println(area)
	fmt.Println(around)
	// var pil int
	// fmt.Println("Luas (1) | Keliling (2) ")
	// fmt.Scan(&pil)
	// switch pil {
	// case 1:
	// 	fmt.Println(area)
	// 	break
	// case 2:
	// 	fmt.Println(around)
	// 	break
	// default:
	// 	os.Exit(0)
	// }
}

// func printPersegi(b bangunDatar) {
// 	luasPersegi := fmt.Sprintf("%.2f", b.lusasBangunDatar())
// 	kelilingPersegi := fmt.Sprintf("%.2f", b.kelilingBangunDatar())
// 	fmt.Println(luasPersegi)
// 	fmt.Println(kelilingPersegi)
// }

func main() {
	hitungLingkaran := lingkaran{7.0}
	hitungPersegi := persegi{5.0}
	printTwoDimentionalFigure(hitungLingkaran)
	printTwoDimentionalFigure(hitungPersegi)
}
