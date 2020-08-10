package power

// Возведение в степень, итеративный метод

import (
	"fmt"
	"time"
)

type Pow1 struct {
}

func (p Pow1) Run(data []string) string {
	a, n, err := parseData(data)

	if err != nil {
		return ""
	}

	t := time.Now()
	r := p.calc(a, n)

	fmt.Println(time.Since(t))

	return floatToString(r)

}

func (p Pow1) calc(a float64, n int) float64 {
	r := float64(1)

	for i := 0; i < n; i++ {
		r *= a
	}

	return r
}
