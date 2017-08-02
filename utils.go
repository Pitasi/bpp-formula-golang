package bppformula

import (
	"math/big"
	"errors"
)

// IntToFloat create a big.Float from a simple int
func IntToFloat(x int) (*big.Float) {
	return new(big.Float).SetInt64(int64(x))
}

// WholePart extracts integer part of a big.Float
func WholePart(number *big.Float) (int64) {
	whole, _ := number.Int64()
	return whole
}

// FractionalPart extracts fractional part of a big.Float
func FractionalPart(number *big.Float) (*big.Float) {
	whole := new(big.Float).SetInt64(WholePart(number))
	frac := new(big.Float).SetPrec(128)

	frac.Sub(number, whole)
	return frac
}

// ToStringBase returns a string representing fractional part of number,
// written in specified base, limited in length to digits
func ToStringBase(number *big.Float, base int, digits int) (string, error) {
	floatBase := IntToFloat(base)
	tmp := new(big.Float).Copy(number)
	res := make([]byte, digits)

	for i := 0; i < digits && tmp.Sign() != 0; i++ {
		frac := FractionalPart(tmp)
		tmp = tmp.Mul(floatBase, frac)
		whole := WholePart(tmp)

		if 0 <= whole && whole <= 9 {
			res[i] = byte('0' + whole)
		} else {
			res[i] = byte('A' + whole - 10)
		}

	}

	return string(res), nil
}

// ModPow calculates (b^n) modulo k, in a memory-efficient way.
func ModPow(b int, n int, k int) (int, error) {
	// TODO: usare meno moltiplicazioni

	if ( k < 1 ) {
		return -1, errors.New("Modulo must be greater or equal than 1")
	}

    if ( k == 1 ) {
		return 0, nil
	}

    b = b % k
    result := 1

    for n > 0 {
        if (n % 2 == 1) {
			result = (result * b) % k
		}
        n = n / 2
        b = (b * b) % k
    }

    return result, nil
}
