package des

func split64bTo32b(block64b uint64) (uint32, uint32) {
	left := uint32(block64b >> 32)
	right := uint32(block64b)
	return left, right
}

func split64bTo8b(block64b uint64) []byte {
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(block64b >> ((7 - i) * 8))
	}
	return bytes
}

func join32bTo64b(left, right uint32) uint64 {
	var block64b uint64
	block64b = uint64(left)
	block64b = (block64b << 32) | uint64(right)
	return block64b
}

func join8bTo64b(blocks8b []byte) uint64 {
	var block64b uint64
	for _, b := range blocks8b {
		block64b = (block64b << 8) | uint64(b)
	}
	return block64b
}
