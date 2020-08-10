package fibonacci

import (
	"fmt"
	"strconv"
	"time"
)

// Расчет чисел Фибоначчи через рекурсию

type Fib1 struct {
}

func (f Fib1) Run(data []string) string {
	n, err := strconv.Atoi(data[0])

	if err != nil {
		return ""
	}

	t := time.Now()
	r := f.calc(n)
	fmt.Println("Execution time", time.Since(t))

	return fmt.Sprint(r)
}

func (f Fib1) calc(n int) uint64 {

	if n == 0 {
		return 0
	}

	if n < 2 {
		return 1
	} else {
		return f.calc(n-1) + f.calc(n-2)
	}
}
