package des

func blockPermutation(block64b uint64, perm [64]byte) uint64 {
	var newBlock64b uint64
	for i := 0; i < 64; i++ {
		newBlock64b |= ((block64b >> (64 - perm[i])) & 1) << (63 - i)
	}
	return newBlock64b
}

func initialPermutation(block64b uint64) uint64 {
	return blockPermutation(block64b, _IP)
}

func finalPermutation(block64b uint64) uint64 {
	return blockPermutation(block64b, _FP)
}
