package calculation

import "interfaceCalculation/utils"

type Pengurangan struct {
	Angka1 int
	Angka2 int
}

func (p Pengurangan) GetCalculation() (int, error) {
	hasilNegatif := p.Angka1 - p.Angka2
	if hasilNegatif < 0 {
		return 0, utils.AngkaNegatifError{"Hasil Pengurangan Negatif"}
	}
	if p.Angka1 == 0 || p.Angka2 == 0 {
		return 0, utils.NumberNotZero{"Salah satu angka ada yang 0"}
	}

	return hasilNegatif, nil
}
