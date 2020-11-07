package des

func initialPermutation(block64b uint64) uint64 {
	var newBlock64b uint64
	for i := 0; i < 64; i++ {
		newBlock64b |= ((block64b >> (64 - _IP[i])) & 1) << (63 - i)
	}
	return newBlock64b
}

func finalPermutation(block64b uint64) uint64 {
	var newBlock64b uint64
	for i := 0; i < 64; i++ {
		newBlock64b |= ((block64b >> (64 - _FP[i])) & 1) << (63 - i)
	}
	return newBlock64b
}
