package des

func bytesAlignment(bytes *[]byte) {
	length := len(*bytes)
	if length%8 != 0 {
		rem := 8 - length%8
		length += rem
		*bytes = append(*bytes, make([]byte, rem)...)
	}
}

func bytesTrim(bytes *[]byte) {
	nullCount := 0
	bytesCount := len(*bytes)
	for i := bytesCount - 1; i >= 0 && (*bytes)[i] == 0; i-- {
		nullCount++
	}
	*bytes = (*bytes)[:bytesCount-nullCount]
}
