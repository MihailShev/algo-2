package power

// Возведение в степень через двоичное разложение показателя степени

import (
	"fmt"
	"time"
)

type Pow3 struct {
}

func (p Pow3) Run(data []string) string {
	a, n, err := parseData(data)

	if err != nil {
		return ""
	}

	t := time.Now()
	r := p.calc(a, n)

	fmt.Println(time.Since(t))

	return floatToString(r)
}

func (p Pow3) calc(a float64, n int) float64 {
	res := float64(1)

	for n > 0 {

		if n%2 == 1 {
			res *= a
		}

		a *= a
		n /= 2
	}

	return res
}
