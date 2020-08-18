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
	if n == 0 {
		return big.NewInt(0)
	}

	if n <= 2 {
		return big.NewInt(1)
	}

	prev := big.NewInt(1)
	curr := big.NewInt(1)

	var result *big.Int

	for i := 2; i < n; i++ {
		result = big.NewInt(0).Add(prev, curr)
		prev = curr
		curr = result
	}

	return curr
}
