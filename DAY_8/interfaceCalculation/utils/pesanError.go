package utils

// AngkaNegatifError Message if result negatif
type AngkaNegatifError struct {
	PesanMasalah string
}

func (err AngkaNegatifError) Error() string {
	return err.PesanMasalah
}

// NumberNotZero Message if any one zero
type NumberNotZero struct {
	MessageBox string
}

func (err NumberNotZero) Error() string {
	return err.MessageBox
}
