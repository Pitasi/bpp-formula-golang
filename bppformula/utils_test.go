package bppformula

import (
	"math/big"
	"testing"
)

func TestEnvFloat(t *testing.T) {
	v49 := new(big.Float).SetInt64(49)
	v1	:= new(big.Float).SetInt64(1)
	res	:= new(big.Float).SetPrec(128)

	res.Quo(v1, v49)				// res = 1/49
	res.Mul(res, v49)				// res = res * 49
	isOne := res.Cmp(v1) == 0		// res == 1

	if res.Prec() != 128 {
		t.Error("Expected to use 128-bit float, instead using", res.Prec())
	}

	if !isOne {
		t.Error("Expected (1/49 * 49 == 1) to be true")
	}
}

func TestModPow(t *testing.T) {
	_, err := ModPow(2, 100, 0)
	if err == nil {
		t.Error("Test an invalid modulo 0, no error were raised")
	}

	test := func(base int, exp int, mod int, expected int64) {
		res, err := ModPow(base, exp, mod)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		if res != expected {
			t.Error("Expected 0, got:", res)
		}
	}

	test(2, 100, 2, 0)
	test(50, 19800, 6, 4)
	test(367, 447, 739, 687)
}

func TestWholePart(t *testing.T) {
	test := func(number float64, expected int64) {
		x := big.NewFloat(number)
		if WholePart(x) != int64(expected) {
			t.Error("Expected", expected, "got", WholePart(x))
		}
	}

	test(456., 456)
	test(1.3456, 1)
	test(423512.23441235, 423512)
}

func TestFractionalPart(t *testing.T) {
	test := func(number float64, expected float64) {
		numFloat := big.NewFloat(number)
		expectedFloat := big.NewFloat(expected)
		frac := FractionalPart(numFloat)
		if expectedFloat.Cmp(frac) != 0 {
			t.Error("Expected", expectedFloat, "got", frac)
		}
	}

	test(0.5, 0.5)
	test(414.25, 0.25)
	test(-16.25, 0.75)
}

func TestToStringBase(t *testing.T) {
	test := func(number float64, base int, expected string) {
		num := big.NewFloat(number)
		str, err := ToStringBase(num, base, 4)
		if err != nil {
			t.Error("Unexpected error:", err)
		}
		if str != expected {
			t.Error("Expected", expected, "got", str)
		}
	}

	test(0.69, 2, "1011")
	test(3.1415926, 16, "243F")
}
