package rsa

import (
	"log"
	"math/big"
)

const (
	_Min          = 0x0001_0000
	_Max          = 0xffff_ffff
	_MillerRabinN = 10
)

type RSA struct {
	e         *big.Int
	d         *big.Int
	n         *big.Int
	blockSize int
}

func getBlockSize(n *big.Int) int {
	nBytes := n.Bytes()
	return len(nBytes)
}

func NewRSA(bits int) *RSA {
	if bits < 0 || bits > 2048 {
		log.Fatal("Bits should be in [1;2048]")
	}
	e, d, n := generateKeys(int64(bits))
	blockSize := getBlockSize(n)

	log.Println("bits =", bits)
	log.Println("E =", e)
	log.Println("D =", d)
	log.Println("N =", n)
	log.Println("blockSize =", blockSize)

	return &RSA{
		e:         e,
		d:         d,
		n:         n,
		blockSize: blockSize,
	}
}

func (rsa *RSA) Encode(text []byte) []byte {
	encText := make([]byte, 0, len(text)*rsa.blockSize)
	for _, m := range text {
		c := rsa.encodeByteToBlock(m)
		encText = append(encText, c...)
	}
	return encText
}

func (rsa *RSA) encodeByteToBlock(m byte) []byte {
	c := make([]byte, 0, rsa.blockSize)
	mBig := big.NewInt(int64(m))
	cBig := new(big.Int).Exp(mBig, rsa.e, rsa.n)
	c = append(c, cBig.Bytes()...)
	dif := rsa.blockSize - len(c)
	if dif != 0 {
		c = append(make([]byte, dif), c...)
	}
	return c
}

func (rsa *RSA) Decode(text []byte) []byte {
	blocks := rsa.splitTextToBlocks(text)
	decText := make([]byte, 0, len(blocks))
	for _, c := range blocks {
		m := rsa.decodeBlockToByte(c)
		decText = append(decText, m)
	}
	return decText
}

func (rsa *RSA) decodeBlockToByte(c []byte) byte {
	cBig := new(big.Int).SetBytes(c)
	mBig := new(big.Int).Exp(cBig, rsa.d, rsa.n)
	m := mBig.Bytes()
	if len(m) == 0 {
		return 0
	}
	if len(m) != 1 {
		log.Fatal("Decoded block is not a byte")
	}
	return m[0]
}

func (rsa *RSA) splitTextToBlocks(text []byte) [][]byte {
	length := len(text)
	blockCount := length / rsa.blockSize
	blocks := make([][]byte, blockCount)
	for i, j := 0, 0; i < blockCount && j < length; i, j = i+1, j+rsa.blockSize {
		blocks[i] = make([]byte, 0, rsa.blockSize)
		blocks[i] = append(blocks[i], text[j:j+rsa.blockSize]...)
	}
	return blocks
}
