package util

import (
	"math/big"

	"github.com/shopspring/decimal"
)

// ZeroAddress is the typical 0x0 in ethereum
const ZeroAddress = "0x0000000000000000000000000000000000000000"

// ToDecimal is a utility function that turns a bigInt into a decimal
func ToDecimal(value *big.Int) decimal.Decimal {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(18)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	return result
}
