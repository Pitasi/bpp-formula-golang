package bppformula

import (
	"math/big"
	"errors"
)

// PRECISION Stop iterating when difference is smaller that this value
// it gives 23-24 exact digits calculating Summations for pi for reasonable d
const PRECISION string = "0.0000000000000000000000000001"

// IntToFloat create a big.Float from a simple int
func IntToFloat(x int) (*big.Float) {
	return Int64ToFloat(int64(x))
}

// Int64ToFloat create a big.Float from a int64
func Int64ToFloat(x int64) (*big.Float) {
	return new(big.Float).SetPrec(128).SetInt64(x)
}

// WholePart extracts integer part of a big.Float
func WholePart(number *big.Float) (int64) {
	whole, _ := number.Int64()
	return whole
}

// Floor performs the floor operation on big.Float, returns a int64
func Floor(number *big.Float) (int64) {
	if number.Sign() >= 0 {
		return WholePart(number)
	}

	return WholePart(number) - 1
}

// FractionalPart extracts fractional part of a big.Float
// defined as x-floor(x), for negative numbers too
func FractionalPart(number *big.Float) (*big.Float) {
	return new(big.Float).Sub(number, Int64ToFloat(Floor(number)))
}

// ToStringBase returns a string representing fractional part of number,
// written in specified base, limited in length to digits
func ToStringBase(number *big.Float, base int, length int) (string, error) {
	if base < 2 {
		return "", errors.New("Base must be greater than or equal to 2")
	}

	floatBase := IntToFloat(base)
	tmp := new(big.Float).Copy(number)
	res := make([]byte, length)

	for i := 0; i < length && tmp.Sign() != 0; i++ {
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

// ModPow calculates (b^n) modulo k, in a memory efficient way
// This is the Right-to-Left binary exponentiation
// For references: Art of Programming vol. 2, par. 4.6.3
func ModPow(b int, n int, k int) (int64, error) {
	if ( k < 1 ) {
		return -1, errors.New("Modulo must be greater than or equal to 1")
	}

    if ( k == 1 ) {
		return 0, nil
	}

    b = b % k
    var result int64 = 1

    for n > 0 {
        if (n % 2 == 1) {
			result = (result * int64(b)) % int64(k)
		}
        n = n / 2
        b = (b * b) % k
    }

    return result, nil
}
