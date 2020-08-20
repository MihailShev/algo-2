package fibonacci

import (
	"fmt"
	"github.com/ALTree/bigfloat"
	"math/big"
)

var ratio = big.NewFloat(1.6180339887498948482045868343656)
var sqr5 = big.NewFloat(2.2360679774997896964091736687313)

type Fib2 struct {
}

func (f Fib2) Run(data []string) string {
	n := big.NewFloat(0)

	n.SetString(data[0])

	tmp := f.calc(n)
	r := big.NewInt(0)
	r, _ = tmp.Int(r)
	return fmt.Sprint(r)
}

func (f Fib2) calc(n *big.Float) *big.Float {
	p := bigfloat.Pow(ratio, n)
	p.Quo(p, sqr5)
	p.Add(p, big.NewFloat(0.5))

	return p
}
