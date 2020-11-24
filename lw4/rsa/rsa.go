package rsa

import (
	"crypto/rand"
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

// TODO: isPrime
func isPrime(n *big.Int) bool {
	return n.ProbablyPrime(_MillerRabinN)
}

func generatePrimeNumber(min, max *big.Int) (n *big.Int) {
	if min == nil {
		min = big.NewInt(_Min)
	}
	if max == nil {
		max = big.NewInt(_Max)
	}
	check := false
	for check != true {
		n, _ = rand.Int(rand.Reader, max)
		if n.Cmp(min) > 0 {
			check = isPrime(n)
		}
	}
	return
}

func eulerFunc(p, q *big.Int) *big.Int {
	one := big.NewInt(1)
	mul1 := new(big.Int).Sub(p, one)
	mul2 := new(big.Int).Sub(q, one)
	return new(big.Int).Mul(mul1, mul2)
}

// TODO: custom GCD
func generateOpenedExp(phi *big.Int) (e *big.Int) {
	minE := big.NewInt(2)
	e = generatePrimeNumber(minE, phi)
	gcdRes := new(big.Int).GCD(nil, nil, e, phi)
	for gcdRes.Cmp(big.NewInt(1)) == 1 {
		e = generatePrimeNumber(minE, phi)
		gcdRes = new(big.Int).GCD(nil, nil, e, phi)
	}
	return
}

// TODO: custom ModInverse
func generateSecretExp(e, phi *big.Int) (d *big.Int) {
	d = new(big.Int).ModInverse(e, phi)
	return
}

func generateP(min, max *big.Int) *big.Int {
	return generatePrimeNumber(min, max)
}

func generateQ(min, max, p *big.Int) *big.Int {
	q := generatePrimeNumber(min, max)
	for p.Cmp(q) == 0 {
		q = generatePrimeNumber(min, max)
	}
	return q
}

func getMin(bits int64) *big.Int {
	x := big.NewInt(2)
	power := big.NewInt(bits - 1)
	min := new(big.Int).Exp(x, power, nil)
	return min
}

func getMax(bits int64) *big.Int {
	x := big.NewInt(2)
	power := big.NewInt(bits)
	max := new(big.Int).Exp(x, power, nil)
	return max
}

func getN(p, q *big.Int) *big.Int {
	return new(big.Int).Mul(p, q)
}

func generateKeys(bits int64) (e, d, n *big.Int) {
	// min = 2**(bits - 1)
	min := getMin(bits)
	// max = 2**bits
	max := getMax(bits)

	// P, Q will be generated as prime numbers in [min; max)
	p := generateP(min, max)
	q := generateQ(min, max, p)
	// n = p*q
	n = getN(p, q)

	// phi = (p - 1)*(q - 1)
	phi := eulerFunc(p, q)
	// GCD(e, phi) = 1
	e = generateOpenedExp(phi)
	// (e*d) % phi = 1
	d = generateSecretExp(e, phi)
	return
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

func (rsa *RSA) Encode(text []byte) []byte {
	encText := make([]byte, 0, len(text)*rsa.blockSize)
	for _, m := range text {
		c := rsa.encodeByteToBlock(m)
		encText = append(encText, c...)
	}
	return encText
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

func (rsa *RSA) Decode(text []byte) []byte {
	blocks := rsa.splitTextToBlocks(text)
	decText := make([]byte, 0, len(blocks))
	for _, c := range blocks {
		m := rsa.decodeBlockToByte(c)
		decText = append(decText, m)
	}
	return decText
}
