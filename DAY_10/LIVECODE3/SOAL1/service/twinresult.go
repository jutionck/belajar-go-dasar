package service

type SliceNumber struct {
	InputSlice []int
	InputParam int
}

func (s SliceNumber) PairsNumber() int {
	var pairsNumber int
	for j, contain := range s.InputSlice {
		for i := 0; i < len(s.InputSlice); i++ {
			if i != j {
				if s.InputSlice[i]+contain == s.InputParam {
					pairsNumber++
					// fmt.Println(contain, inputSlice[i])
					s.InputSlice[j] = 0
					s.InputSlice[i] = 0
				}
			}
		}
	}
	return pairsNumber
}
