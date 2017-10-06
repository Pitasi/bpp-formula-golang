package bbpformula

import (
  "testing"
)

func TestS(t *testing.T) {
  // Single tests for partial summations
  // Expected values taken from:
  // http://www.davidhbailey.com/dhbpapers/bbp-alg.pdf, page 4

  test := func(j int, d int, expected string) {
    calc := S(1, j, d).Text('f', 30)[0:30] // truncate at 28 decimal digits
    if calc != expected {
      t.Error("Expected", expected, "got", calc)
    }
  }

  d := 1000000
  test(1, d, "0.1810395338014360678534893462")
  test(4, d, "0.7760655498078074613722975943")
  test(5, d, "0.3624585640705741420683343355")
  test(6, d, "0.3861386739520148480012151865")
}

func TestPi(t *testing.T) {
  test := func(d int, expected string) {
    calc, err := ToStringBase(Pi(d), 16, 8)
    if err != nil {
      t.Error("Unexpected error", err)
    }
    if calc != expected {
      t.Error("Expected", expected, "got", calc)
    }
  }

  test(1, "243F6A88")
  test(1000001, "6C65E52C")
}

func TestPiOverlap(t *testing.T) {
  test := func(d int) {
    first, _ := ToStringBase(Pi(d), 16, 8)
    second, _ := ToStringBase(Pi(d+4), 16, 8)
    if first[4:8] != second[0:4] {
      t.Error("Digits don't overlap. Expected", first[4:8],
        "and", second[0:4], "to be equal")
    }
  }

  test(1)
  test(10)
  test(12323)
}
