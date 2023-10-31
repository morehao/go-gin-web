package utils

import (
	"fmt"
	"strconv"
)

func SetPrec(num float64, precision int) float64 {
	precisionRule := fmt.Sprintf("%%.%df", precision)
	rate, _ := strconv.ParseFloat(fmt.Sprintf(precisionRule, num), 64)
	return rate
}
