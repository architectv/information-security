package des

func (des *DES) keyPermutation56bTo28b() (uint32, uint32) {
	var k1, k2 uint64
	for i := 0; i < 28; i++ {
		k1 |= ((des.key >> (64 - _K1P[i])) & 1) << (31 - i)
		k2 |= ((des.key >> (64 - _K2P[i])) & 1) << (31 - i)
	}
	return uint32(k1), uint32(k2)
}

func (des *DES) keyExpansion() [_RoundCount]uint64 {
	var keys48b [_RoundCount]uint64
	k1, k2 := des.keyPermutation56bTo28b()

	var n byte
	var block56b uint64
	lShift28b := func(x uint32, n byte) uint32 {
		return ((x << n) | (x >> (-n & 27))) & ((1 << 32) - 1)
	}
	for i := 0; i < _RoundCount; i++ {
		switch i {
		case 0, 1, 8, 15:
			n = 1
		default:
			n = 2
		}

		k1 = lShift28b(k1, n)
		k2 = lShift28b(k2, n)

		block56b = uint64(k1 >> 4)
		block56b = ((block56b << 32) | uint64(k2)) << 4

		for j := 0; j < 48; j++ {
			keys48b[i] |= ((block56b >> (64 - _CP[j])) & 1) << (63 - j)
		}
	}

	return keys48b
}
