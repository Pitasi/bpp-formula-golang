package bppformula

import (
  "math"
  "math/big"
)

// Log2 computes log2 at binary digit specified, returns a big.Float
func Log2(digit int) (*big.Float) {
  digit--
  k := 1
  res := big.NewFloat(0).SetPrec(128)

  // First summation (from 1 to d)
  for k <= digit {
    numerator, err := ModPow(2, digit-k, k)
    if err != nil {
      // TODO: handle err
    }
    tmpRes := new(big.Float).SetPrec(128).
      Quo(Int64ToFloat(numerator), IntToFloat(k)) // quotient = modpow/k
    res.Add(res, tmpRes)
    k++
  }
  res = FractionalPart(res)

  // Seconds summation (from d+1 to infinite/precision)
  floatPrecision, _ := new(big.Float).SetString(PRECISION)
  delta := IntToFloat(1)
  for delta.Cmp(floatPrecision) >= 0 {
    denominator := IntToFloat(k)
    numerator := big.NewFloat(math.Pow(2, float64(digit-k)))
    delta.Quo(numerator, denominator)
    res.Add(res, delta)
    k++
  }

  return res
}
