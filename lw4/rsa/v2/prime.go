package rsa

import (
	"crypto/rand"
	"math/big"
)

// v2: [isPrime] customProbablyPrime
func isPrime(n *big.Int) bool {
	return customProbablyPrime(n, _MillerRabinN)
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

// v2: [customProbablyPrime] Miller-Rabin Test
func customProbablyPrime(n *big.Int, k int) bool {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)
	if n.Cmp(two) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}
	if n.Cmp(two) < 0 || new(big.Int).Mod(n, two).Cmp(zero) == 0 {
		return false
	}
	t := new(big.Int).Sub(n, one)
	s := 0
	for new(big.Int).Mod(t, two).Cmp(zero) == 0 {
		t = new(big.Int).Div(t, two)
		s++
	}
	for i := 0; i < k; i++ {
		max := new(big.Int).Sub(n, one)
		a := one
		for a.Cmp(two) < 0 {
			a, _ = rand.Int(rand.Reader, max)
		}

		x := customExp(a, t, n)
		if x.Cmp(one) == 0 || x.Cmp(max) == 0 {
			continue
		}

		for r := 1; r < s; r++ {
			x = customExp(x, two, n)
			if x.Cmp(one) == 0 {
				return false
			}
			if x.Cmp(max) == 0 {
				break
			}
		}

		if x.Cmp(max) != 0 {
			return false
		}
	}

	return true
}
