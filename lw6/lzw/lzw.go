package lzw

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Bits int64 = 256
)

func Compress(data []byte) []byte {
	dictSize := Bits
	dictionary := make(map[string]int64, dictSize)
	for i := int64(0); i < dictSize; i++ {
		dictionary[string([]byte{byte(i)})] = i
	}

	var result []int64
	var w []byte
	for i := 0; i < len(data); i++ {
		c := data[i]
		wc := append(w, c)
		if _, ok := dictionary[string(wc)]; ok {
			w = wc
		} else {
			result = append(result, dictionary[string(w)])
			dictionary[string(wc)] = dictSize
			dictSize++
			wc[0] = c
			w = wc[:1]
		}
	}

	if len(w) > 0 {
		result = append(result, dictionary[string(w)])
	}

	return toBytes(result)
}

func Decompress(data []byte) []byte {
	compressed := fromBytes(data)
	dictSize := Bits
	dictionary := make(map[int64][]byte, dictSize)
	for i := int64(0); i < dictSize; i++ {
		dictionary[i] = []byte{byte(i)}
	}

	var result strings.Builder
	var w []byte
	for _, k := range compressed {
		var entry []byte
		if x, ok := dictionary[k]; ok {
			entry = x[:len(x):len(x)]
		} else if k == dictSize && len(w) > 0 {
			entry = append(w, w[0])
		} else {
			return []byte(result.String())
		}
		result.Write(entry)

		if len(w) > 0 {
			w = append(w, entry[0])
			dictionary[dictSize] = w
			dictSize++
		}
		w = entry
	}

	return []byte(result.String())
}

func toBytes(data []int64) []byte {
	binStr := ""
	bits := 9
	idx := Bits
	for _, v := range data {
		if idx > 1<<bits {
			bits++
		}
		format := "%0" + strconv.Itoa(bits) + "b"
		binStr += fmt.Sprintf(format, v)
		idx++
	}

	bytes := make([]byte, 0)
	length := len(binStr)
	size := 8
	for i := 0; i < length; i += size {
		end := i + size
		if end > length {
			end = length
		}
		binElem := binStr[i:end]
		elem, err := strconv.ParseInt(binElem, 2, size+1)
		if err != nil {
			fmt.Println(err.Error())
		}
		bytes = append(bytes, byte(elem))
	}
	return bytes
}

func trimBinString(s string) string {
	length := len(s)
	trim := ""
	idx := Bits
	bits := 9
	i := 0
	for ; i < length; i += bits {
		if i+bits > length {
			break
		}
		if idx > 1<<bits {
			bits++
		}
		idx++
	}

	size := 8
	mod := (length - i) % size
	get := size - mod
	elem := s[length-get : length]
	trim = s[:length-size] + elem

	return trim
}

func fromBytes(data []byte) []int64 {
	binStr := ""
	for _, b := range data {
		binElem := fmt.Sprintf("%08b", b)
		binStr += binElem
	}

	binStr = trimBinString(binStr)

	compressed := make([]int64, 0)
	bits := 9
	length := len(binStr)
	idx := 256
	for i := 0; i < length; i += bits {
		if idx > 1<<bits {
			bits++
		}
		start := i
		end := i + bits
		if end > length {
			end = length
			start = end - bits
		}
		binElem := binStr[start:end]
		elem, err := strconv.ParseInt(binElem, 2, bits+1)
		if err != nil {
			fmt.Println(err.Error())
		}
		compressed = append(compressed, elem)
		idx++
	}
	return compressed
}
