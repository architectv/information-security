package des

func feistelCipherMode(mode byte) (
	feistelCipher func(*uint32, *uint32, [_RoundCount]uint64)) {
	switch mode {
	case 'e':
		feistelCipher = feistelCipherEnc
	case 'd':
		feistelCipher = feistelCipherDec
	}
	return
}

func feistelCipherEnc(left, right *uint32, keys48b [_RoundCount]uint64) {
	for round := 0; round < _RoundCount; round++ {
		roundFeistelCipher(left, right, keys48b[round])
	}
	*left, *right = *right, *left
}

func feistelCipherDec(left, right *uint32, keys48b [_RoundCount]uint64) {
	for round := 15; round >= 0; round-- {
		roundFeistelCipher(left, right, keys48b[round])
	}
	*left, *right = *right, *left
}

func roundFeistelCipher(left, right *uint32, key48b uint64) {
	tmp := *right
	*right = funcF(*right, key48b) ^ *left
	*left = tmp
}
