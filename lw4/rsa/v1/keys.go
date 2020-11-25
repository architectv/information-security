package rsa

import (
	"log"
	"math/big"
)

func eulerFunc(p, q *big.Int) *big.Int {
	one := big.NewInt(1)
	mul1 := new(big.Int).Sub(p, one)
	mul2 := new(big.Int).Sub(q, one)
	return new(big.Int).Mul(mul1, mul2)
}

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
	log.Println("P =", p)
	log.Println("Q =", q)
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
