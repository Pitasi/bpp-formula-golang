package bppformula

import (
	"math"
	"math/big"
)

// PRECISION Stop iterating when difference is more little that this value
const PRECISION float64 = 0.00000001

// Log2 calculates log2 at binary digit specified, returns a big.Float
func Log2(digit int) (*big.Float) {
	digit--
	k := 1
	res := big.NewFloat(0).SetPrec(128)

	for k <= digit {
		modpow, err := ModPow(2, digit-k, k)
		if (err != nil) {
			// TODO: handle err
		}
		quotient := new(big.Float).SetPrec(128).
			Quo(IntToFloat(modpow), IntToFloat(k)) // quotient = modpow/k
		res.Add(res, quotient)
		k++
	}
	res = FractionalPart(res)

	floatPrecision := big.NewFloat(PRECISION)
	diff := IntToFloat(1)
	for diff.Cmp(floatPrecision) >= 0 {
		floatK := IntToFloat(k)
		pow := math.Pow(2, float64(digit-k))
		floatPow := big.NewFloat(pow)
		diff.Quo(floatPow, floatK)
		res.Add(res, diff)
		k++
	}

	return res
}
