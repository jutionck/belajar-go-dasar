package pengurangan

import "DICTIONARYINTERFACE/utils"

// Pengurangan struct
type Pengurangan struct {
	Angka1 int
	Angka2 int
}

func (p Pengurangan) GetPengurangan() (int, error) {
	hasilNegatif := p.Angka1 - p.Angka2
	if hasilNegatif < 0 {
		return 0, utils.DibagiNolError{"Hasil pengurangan negatif"}
	} else {
		return p.Angka1 - p.Angka2, nil
	}
}
