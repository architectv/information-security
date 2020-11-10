package des

type DES struct {
	key uint64
}

func NewDES(keyString string) *DES {
	key := join8bTo64b([]byte(keyString))
	return &DES{
		key: key,
	}
}

func (des *DES) Encode(text []byte) []byte {
	length := len(text)
	lengthBytes := split64bTo8b(uint64(length))
	encLengthBytes := des.code(_EncMode, lengthBytes)

	encText := des.code(_EncMode, text)
	encText = append(encText, encLengthBytes...)
	return encText
}

func (des *DES) Decode(text []byte) []byte {
	length := len(text)
	idxLastBlock := length - 8
	lengthBlock := text[idxLastBlock:]
	decLengthBlock := des.code(_DecMode, lengthBlock)
	decLength := int(join8bTo64b(decLengthBlock))

	decText := des.code(_DecMode, text[:idxLastBlock])
	decText = decText[:decLength]
	return decText
}

func (des *DES) code(mode byte, text []byte) []byte {
	feistelCipher := feistelCipherMode(mode)
	bytesAlignment(&text)
	length := len(text)

	encText := make([]byte, 0, length)
	keys48b := des.keyExpansion()

	for i := 0; i < length; i += 8 {
		block64b := join8bTo64b(text[i : i+8])
		ip := initialPermutation(block64b)
		left, right := split64bTo32b(ip)

		feistelCipher(&left, &right, keys48b)

		newBlock64b := join32bTo64b(left, right)
		fp := finalPermutation(newBlock64b)
		encBlock := split64bTo8b(fp)
		encText = append(encText, encBlock...)
	}

	return encText
}

func (des *DES) GetKey() uint64 {
	return des.key
}
