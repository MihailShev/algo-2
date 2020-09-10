package fibonacci

// Расчет чисел Фибоначчи через рекурсию

import (
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

type Fib1 struct {
}

func (f Fib1) Run(data []string) string {
	n := big.NewInt(0)

	n.SetString(data[0], 10)

	r := f.calc(n)

	return fmt.Sprint(r)
}

func (f Fib1) calc(n *big.Int) *big.Int {

	if n.Cmp(zero) == 0 {
		return zero
	}

	if n.Cmp(big.NewInt(2)) == -1 || n.Cmp(big.NewInt(2)) == 0 {
		return one
	}

	a := big.NewInt(0).Sub(n, one)
	b := big.NewInt(0).Sub(n, two)

	resA := f.calc(a)
	resB := f.calc(b)

	return big.NewInt(0).Add(resA, resB)

}
