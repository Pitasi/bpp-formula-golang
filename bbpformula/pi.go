package bbpformula

import (
  "math"
  "math/big"
)

// Pi computes pi at hexadecimal digit specified, returns a big.Float
func Pi(digit int) (*big.Float) {
  digit--
  res := S(4, 1, digit)
  res.Sub(res, S(2, 4, digit))
  res.Sub(res, S(1, 5, digit))
  res.Sub(res, S(1, 6, digit))

  return FractionalPart(res)
}

// S computes x*Summation(j,d)
func S(x int, j int, d int) (*big.Float) {
  sum := Summation(j, d)
  if x == 1 {
    return sum
  }
  return sum.Mul(sum, IntToFloat(x))
}

// Summation computes {16^d*Sj}, see PDF attached for references
func Summation(j int, d int) (*big.Float) {
  k := 0
  res := IntToFloat(0)

  // First summation (from 1 to d)
  for k <= d {
    denominator := (8*k)+j
    numerator, err := ModPow(16, d-k, denominator)
    if err != nil {
      // TODO: handle error
    }
    tmpRes := new(big.Float).SetPrec(128).
          Quo(Int64ToFloat(numerator), IntToFloat(denominator))
    res.Add(res, tmpRes)
    k++
  }
  res = FractionalPart(res)

  // Seconds summation (from d+1 to infinite/precision)
  floatPrecision, _ := new(big.Float).SetPrec(128).SetString(PRECISION)
  delta := IntToFloat(1)
  for delta.Cmp(floatPrecision) >= 0 {
    denominator := (8*k)+j
    numerator := math.Pow(16, float64(d-k))
    delta.Quo(big.NewFloat(numerator), IntToFloat(denominator))
    res.Add(res, delta)
    k++
  }

  return FractionalPart(res)
}
