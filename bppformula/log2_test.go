package bppformula

import (
	"testing"
)

func TestLog2(t *testing.T) {
	test := func(d int, expected string) {
		calc, _ := ToStringBase(Log2(d), 2, 8)
		if (calc != expected) {
			t.Error("Expected", expected, "got", calc)
		}
	}

	test(1, "10110001")
}

func TestLog2Overlap(t *testing.T) {
	test := func(d int) {
		first, _ := ToStringBase(Log2(d), 2, 8)
		second, _ := ToStringBase(Log2(d+4), 2, 8)
		if (first[4:8] != second[0:4]) {
			t.Error("Digits don't overlap. Expected", first[4:8],
				"and", second[0:4], "to be equal")
		}
	}

	test(1)
	test(1000)
	test(1234143)
}
