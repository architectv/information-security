package rsa

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomGCD(t *testing.T) {
	tests := []struct {
		name string
		a    *big.Int
		b    *big.Int
		gcd  *big.Int
	}{
		{
			name: "GCD should be 1",
			a:    big.NewInt(3),
			b:    big.NewInt(5),
			gcd:  big.NewInt(1),
		},
		{
			name: "GCD should NOT be 1",
			a:    big.NewInt(6),
			b:    big.NewInt(12),
			gcd:  big.NewInt(6),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := customGCD(test.a, test.b)
			assert.Equal(t, test.gcd, res)
		})
	}
}

func TestCustomExp(t *testing.T) {
	tests := []struct {
		name string
		x    *big.Int
		y    *big.Int
		z    *big.Int
		exp  *big.Int
	}{
		{
			name: "Exp #1",
			x:    big.NewInt(5),
			y:    big.NewInt(2),
			z:    big.NewInt(4),
			exp:  big.NewInt(1),
		},
		{
			name: "Exp #2",
			x:    big.NewInt(5),
			y:    big.NewInt(100),
			z:    big.NewInt(7),
			exp:  big.NewInt(2),
		},
		{
			name: "Exp #3",
			x:    big.NewInt(3),
			y:    big.NewInt(90),
			z:    big.NewInt(5),
			exp:  big.NewInt(4),
		},
		{
			name: "Exp #4",
			x:    big.NewInt(49),
			y:    big.NewInt(6431),
			z:    big.NewInt(35),
			exp:  big.NewInt(14),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := customExp(test.x, test.y, test.z)
			assert.Equal(t, test.exp, res)
		})
	}
}

func TestCustomProbablyPrime(t *testing.T) {
	const Rounds = _MillerRabinN
	tests := []struct {
		name    string
		number  *big.Int
		rounds  int
		isPrime bool
	}{
		{
			name:    "ProbablyPrime #1",
			number:  big.NewInt(2),
			rounds:  Rounds,
			isPrime: true,
		},
		{
			name:    "ProbablyPrime #2",
			number:  big.NewInt(5),
			rounds:  Rounds,
			isPrime: true,
		},
		{
			name:    "ProbablyPrime #3",
			number:  big.NewInt(8),
			rounds:  Rounds,
			isPrime: false,
		},
		{
			name:    "ProbablyPrime #4",
			number:  big.NewInt(13),
			rounds:  Rounds,
			isPrime: true,
		},
		{
			name:    "ProbablyPrime #5",
			number:  big.NewInt(36),
			rounds:  Rounds,
			isPrime: false,
		},
		{
			name:    "ProbablyPrime #6",
			number:  big.NewInt(71),
			rounds:  Rounds,
			isPrime: true,
		},
		{
			name:    "ProbablyPrime #7",
			number:  big.NewInt(839),
			rounds:  Rounds,
			isPrime: true,
		},
		{
			name:    "ProbablyPrime #8",
			number:  big.NewInt(695531), // 829*839
			rounds:  Rounds,
			isPrime: false,
		},
		{
			name:    "ProbablyPrime #9",
			number:  big.NewInt(988939464559), // 9973*9967*9949
			rounds:  Rounds,
			isPrime: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := customProbablyPrime(test.number, test.rounds)
			assert.Equal(t, test.isPrime, res)
		})
	}
}

func TestCustomModInverse(t *testing.T) {
	const Rounds = _MillerRabinN
	tests := []struct {
		name string
		a    *big.Int
		m    *big.Int
		want *big.Int
	}{
		{
			name: "ModInverse #1",
			a:    big.NewInt(3),
			m:    big.NewInt(11),
			want: big.NewInt(4),
		},
		{
			name: "ModInverse #2",
			a:    big.NewInt(10),
			m:    big.NewInt(17),
			want: big.NewInt(12),
		},
		{
			name: "ModInverse #3",
			a:    big.NewInt(3),
			m:    big.NewInt(9167368),
			want: big.NewInt(6111579),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := customModInverse(test.a, test.m)
			assert.Equal(t, test.want, res)
		})
	}
}
