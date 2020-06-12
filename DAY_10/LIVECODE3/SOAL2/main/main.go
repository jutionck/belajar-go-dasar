/*
Seorang penjudi memiliki Kredit Balance di Slot Machine sebanyak 5000$ untuk melakukan pemutaran pada mesin Slot Machine. Slot Machine adalah sebuah permainan judi kasino yang akan menghasilkan tiga angka ( range angka hanya 1-9 ).

1. Apabila mendapatkan 3 angka dengan nilai yang berbeda, maka kredit balance akan dikurangi dari setiap angka dikalikan dengan 50. Contoh hasil Slot Machine adalah 1, 2, dan 3. Maka Kredit Balance akan dikurangi sebesar (1 + 2 + 3) * 50 = 300
2. Apabila mendapatkan 2 angka yang sama, maka kamu akan mendapat kredit sebesar total dari  setiap angka dikalikan dengan 100. Contoh hasil Slot Machine adalah 1, 1, dan 2. Maka Kredit Balance akan ditambahkan sebesar (1 + 1 + 2) * 100 = 400
3. Apabila mendapatkan 3 angka yang sama, maka kamu akan mendapat kredit sebesar total dari setiap angka dikalikan dengan 200. Contoh hasil Slot Machine adalah 1, 1, 1. Maka Kredit Balance akan ditambahkan sebesar (1 + 1 + 1) * 200 = 600

Buatlah sebuah function yang akan mengembalikan string berupa

1. "You Win X Dollars" apabila tidak kalah, dimana X merupakan kredit yang dimenangkan penjudi dan "Your Total Credit Balance Is X Dollar" merupakan hasil dari penambahan kredit awal + kredit yang dimenangkan
2. "You Lose X Dollars" apabila kalah (mendapatkan tiga angka berbeda dari Slot Machine), dimana X merupakan kredit yang dibayarkan penjudi dan "Your Total Credit Balance Is X Dollar" merupakan hasil dari pengurangan kredit awal - kredit yang dibayarkan

Code By : Jution Candra Kirana
*/

package main

import (
	"LIVECODE3/SOAL2/service"
	"fmt"
)

func printResult(r service.Result) {
	hasil := r.Message()
	fmt.Println(hasil)
}

func main() {
	var deret = []int{1, 1, 2}
	creditBalance := 5000
	resultFinal := service.SlotMachine{deret, creditBalance}
	printResult(resultFinal)
}
