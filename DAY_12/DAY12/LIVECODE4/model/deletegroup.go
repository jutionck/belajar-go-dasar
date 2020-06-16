package model

import (
	f "fmt"
)

func deleteGroup() {
	if totalGroup() < 1 {
		f.Println("Error: data group masih kosong")
	} else {
		i := totalGroup()
		groups = groups[:i-1]
		f.Println()
		f.Println("Success: data group yang terakhir masuk berhasil dihapus")
		f.Println(groups)
		f.Println()
		MainMenu()
	}
}
