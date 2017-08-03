package bppformula

import (
	"math/big"
	"testing"
)

func TestEnvFloat(t *testing.T)() {
	v49 := new(big.Float).SetInt64(49)
	v1	:= new(big.Float).SetInt64(1)
	res	:= new(big.Float).SetPrec(128)

	res.Quo(v1, v49)				// res = 1/49
	res.Mul(res, v49)				// res = res * 49
	isOne := res.Cmp(v1) == 0		// res == 1

	if (res.Prec() != 128) {
		t.Error("Expected to use 128-bit float, instead using", res.Prec())
	}

	if (!isOne) {
		t.Error("Expected (1/49 * 49 == 1) to be true")
	}
}

func TestModPow(t *testing.T)() {
	_, err := ModPow(2, 100, 0)
	if (err == nil) {
		t.Error("Invalid modulo 0, expected to return error")
	}

	res, err := ModPow(2, 100, 2)
	if (err != nil) {
		t.Error("Unexpected error:", err)
	}
	if (res != 0) {
		t.Error("Expected 0, got:", res)
	}

	res, err = ModPow(50, 19800, 6)
	if (err != nil) {
		t.Error("Unexpected error:", err)
	}
	if (res != 4) {
		t.Error("Expected 4, got:", res)
	}
}

func TestWholePart(t *testing.T)() {
	x := new(big.Float).SetFloat64(456)
	if (WholePart(x) != int64(456)) {
		t.Error("Expected", 456, "got", WholePart(x))
	}

	x = new(big.Float).SetFloat64(1.3456)
	var whole int64 = 1
	if (WholePart(x) != whole) {
		t.Error("Expected", whole, "got", WholePart(x))
	}
}

func TestFractionalPart(t *testing.T)() {
	x := new(big.Float).SetFloat64(0.5)
	expected := new(big.Float).SetFloat64(0.5)
	frac := FractionalPart(x)
	if (expected.Cmp(frac) != 0) {
		t.Error("Expected", expected, "got", frac)
	}

	x = new(big.Float).SetFloat64(414.25)
	expected = new(big.Float).SetFloat64(0.25)
	frac = FractionalPart(x)
	if (expected.Cmp(frac) != 0) {
		t.Error("Expected", expected, "got", frac)
	}

	x = new(big.Float).SetFloat64(-16.25)
	expected = new(big.Float).SetFloat64(0.75)
	frac = FractionalPart(x)
	if (expected.Cmp(frac) != 0) {
		t.Error("Expected", expected, "got", frac)
	}
}

func TestToStringBase(t *testing.T)() {
	num := new(big.Float).SetFloat64(0.69)
	str, err := ToStringBase(num, 2, 4)
	if (err != nil) {
		t.Error("Unexpected error:", err)
	}
	if (str != "1011") {
		t.Error("Expected 1011, got", str)
	}
}
