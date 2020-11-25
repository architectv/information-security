package rsa

import (
	"math/big"
)

func customExtendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if a.Cmp(b) > 0 {
		a, b = b, a
	}
	prevX := big.NewInt(1)
	x := big.NewInt(0)
	prevY := big.NewInt(0)
	y := big.NewInt(1)

	for b.Cmp(big.NewInt(0)) != 0 {
		q := new(big.Int).Div(a, b)

		newX := new(big.Int).Sub(prevX, new(big.Int).Mul(q, x))
		x, prevX = newX, x

		newY := new(big.Int).Sub(prevY, new(big.Int).Mul(q, y))
		y, prevY = newY, y

		a, b = b, new(big.Int).Mod(a, b)
	}

	return a, prevX, prevY
}

func customModInverse(a, m *big.Int) *big.Int {
	g, x, _ := customExtendedGCD(a, m)
	if g.Cmp(big.NewInt(1)) != 0 {
		return nil
	}
	return new(big.Int).Mod(x, m)
}

func customGCD(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		a, b = b, a
	}
	for b.Cmp(big.NewInt(0)) != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

func customExp(x, y, z *big.Int) *big.Int {
	n := y.BitLen()
	res := big.NewInt(1)
	power := big.NewInt(2)
	for i := 0; i < n; i++ {
		if i != 0 {
			x = new(big.Int).Exp(x, power, z)
		}
		bit := y.Bit(i)
		bitPower := big.NewInt(int64(bit))
		factor := new(big.Int).Exp(x, bitPower, nil)
		res = new(big.Int).Mul(res, factor)
	}
	res = new(big.Int).Mod(res, z)
	return res
}
