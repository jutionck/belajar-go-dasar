package utils

// DibagiNolError is User Defined Error
type DibagiNolError struct {
	PesanMasalah string
}
type AngkaMinus struct {
	PesanMasalah string
}

func (dne DibagiNolError) Error() string {
	if len(dne.PesanMasalah) == 0 {
		return DIVIDED_BY_ZERO
	}
	return dne.PesanMasalah
}
func (dne AngkaMinus) Error() string {
	if len(dne.PesanMasalah) == 0 {
		return MINUS
	}
	return dne.PesanMasalah
}
