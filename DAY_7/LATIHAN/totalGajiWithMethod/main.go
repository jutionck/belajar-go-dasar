/*
Contoh program sederhana menampilkan :
1. totalGaji karyawan (full)
2. totalGajiBulanan karyawan
By : Bang Edward
*/
package main

import (
	"fmt"
)

type gaji struct {
	basic, tunjangan, lembur float64
}

type karyawan struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []gaji
}

func (e *karyawan) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.basic)
		fmt.Println(info.tunjangan)
		fmt.Println(info.lembur)
		fmt.Println("Sub Total :", info.basic+info.lembur+info.tunjangan)
		// fmt.Println("Harian :", ((5/20)*info.basic + info.lembur + info.tunjangan))

	}
	return "----------------------"
}

func (e *karyawan) sumSallary() (totalGajinya float64) {

	for _, info := range e.MonthlySalary {
		totalGajinya += (info.basic + info.lembur + info.tunjangan)
	}
	return
}

func (e *karyawan) sumSallaryDaily() (totalGajiHarian float64) {
	/*
		(jumlah jam kerja / jumlah jam kerja sebulan) * gaji sebulan
	*/
	for _, info := range e.MonthlySalary {
		totalGajiHarian = (15 / 60) * (totalGajiHarian + (info.basic + info.lembur + info.tunjangan))
	}
	return
}

func main() {
	e := karyawan{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []gaji{
			gaji{basic: 15000.00, tunjangan: 5000.00, lembur: 2000.00},
			gaji{basic: 16000.00, tunjangan: 5000.00, lembur: 2100.00},
			gaji{basic: 17000.00, tunjangan: 5000.00, lembur: 2200.00},
		},
	}
	fmt.Println(e.EmpInfo())
	f := karyawan{
		MonthlySalary: []gaji{
			gaji{basic: 15000.00, tunjangan: 5000.00, lembur: 2000.00},
			gaji{basic: 16000.00, tunjangan: 5000.00, lembur: 2100.00},
			gaji{basic: 17000.00, tunjangan: 5000.00, lembur: 2200.00},
		},
	}
	fmt.Println("Total Gajinya :")
	fmt.Println(f.sumSallary()) //total gaji
	fmt.Println("Gaji Perhari :")
	fmt.Println(f.sumSallaryDaily()) //gaji harian
}
