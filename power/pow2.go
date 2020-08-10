package power

// Возведение в степень через степень двойки с домножением

import (
	"fmt"
	"time"
)

type Pow2 struct {
}

func (p Pow2) Run(data []string) string {
	a, n, err := parseData(data)

	if err != nil {
		return ""
	}

	t := time.Now()
	r := p.calc(a, n)

	fmt.Println(time.Since(t))

	return floatToString(r)
}

func (p Pow2) calc(a float64, n int) float64 {
	if n < 4 {
		return Pow1{}.calc(a, n)
	}

	r := a * a

	usePow2 := true

	for i := 4; i <= n; {
		if usePow2 {
			r = r * r

			if i+i > n {
				usePow2 = false
				i++
			} else {
				i += i
			}
		} else {
			r *= a
			i++
		}
	}

	return r
}
