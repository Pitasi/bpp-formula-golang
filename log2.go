package bppformula

import (
	"math"
	"math/big"
)

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
			Quo(Int64ToFloat(modpow), IntToFloat(k)) // quotient = modpow/k
		res.Add(res, quotient)
		k++
	}
	res = FractionalPart(res)

	floatPrecision, _ := new(big.Float).SetString(PRECISION)
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
