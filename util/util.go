package util

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func ToDecimal(value *big.Int) decimal.Decimal {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(18)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	return result
}
