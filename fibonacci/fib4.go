package fibonacci

import (
	"fmt"
	"math/big"
	"strconv"
)

// Расчет чисел Фибоначчи итеративным способом

type Fib4 struct {
}

func (f Fib4) Run(data []string) string {
	n, err := strconv.Atoi(data[0])

	if err != nil {
		return ""
	}

	r := f.calc(n)

	return fmt.Sprint(r)
}

func (f Fib4) calc(n int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		a, b = b, big.NewInt(0).Add(a, b)
	}

	return a
}
