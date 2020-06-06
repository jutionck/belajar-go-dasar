// Menggunakan Switch Case
package main

import (
	"fmt"
)

func main() {
	// input
	var tanggal int = 21
	var bulan int = 2
	var tahun int = 2020
	var konvbulan string
	switch bulan {
	case 1:
		konvbulan = "Januari"
	case 2:
		konvbulan = "Februari"
	case 3:
		konvbulan = "Maret"
	case 4:
		konvbulan = "April"
	case 5:
		konvbulan = "Mei"
	case 6:
		konvbulan = "Juni"
	case 7:
		konvbulan = "Juli"
	case 8:
		konvbulan = "Agustus"
	case 9:
		konvbulan = "September"
	case 10:
		konvbulan = "Oktober"
	case 11:
		konvbulan = "November"
	case 12:
		konvbulan = "Desember"
	}

	fmt.Println(tanggal, konvbulan, tahun)
}
