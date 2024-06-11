package helpers

type Word []byte

func NewWordFrom(word string) Word {
	return Word([]byte(word))
}

func (w Word) sum() uint64 {
	var result uint64

	for _, b := range w {
		result = result + uint64(b)
	}

	return result
}

type Words []Word

func (w Words) sum() uint64 {
	var result uint64

	for _, word := range w {
		result = result + word.sum()
	}

	return result
}

func (w Word) Hash(salts ...Word) uint64 {
	var result uint64

	for _, b := range w {
		result = (result << 5) + uint64(b) + Words(salts).sum()
	}

	return result
}
