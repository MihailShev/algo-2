package fibonacci

import (
	"fmt"
	"math"
	"strconv"
)

const ratio = 1.6180339887498948482045868343656
const sqr5 = 2.2360679774997896964091736687313

type Fib2 struct {
}

func (f Fib2) Run(data []string) string {
	n, err := strconv.Atoi(data[0])

	if err != nil {
		return ""
	}

	r := f.calc(n)

	return fmt.Sprint(r)
}

func (f Fib2) calc(n int) uint64 {
	r := math.Pow(ratio, float64(n))/sqr5 + 0.5
	return uint64(r)
}
