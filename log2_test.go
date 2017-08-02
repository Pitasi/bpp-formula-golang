package bppformula

import (
	"testing"
)

func TestLog2(t *testing.T)() {
	res, _ := ToStringBase(Log2(1), 2, 8)
	if (res != "10110001") {
		t.Error("Expected 10110001, got", res)
	}

	first, _ := ToStringBase(Log2(10000), 2, 8)
	second, _ := ToStringBase(Log2(10004), 2, 8)
	if (first[4:8] != second[0:4]) {
		t.Error("Digits don't overlap. Expected", first[4:8],
			"and", second[0:4], "to be equal")
	}

	first, _ = ToStringBase(Log2(1234143), 2, 8)
	second, _ = ToStringBase(Log2(1234147), 2, 8)
	if (first[4:8] != second[0:4]) {
		t.Error("Digits don't overlap. Expected", first[4:8],
			"and", second[0:4], "to be equal")
	}
}
