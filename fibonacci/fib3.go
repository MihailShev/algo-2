package fibonacci

// Расчет чисел Фибоначчи через возведение матрицы в степень

import (
	"fmt"
	"strconv"
)

var matrix = [2][2]uint64{[2]uint64{1, 1}, [2]uint64{1, 0}}

type Fib3 struct {
}

func (f Fib3) Run(data []string) string {
	n, err := strconv.Atoi(data[0])

	if err != nil {
		return ""
	}

	r := f.calc(n)

	return fmt.Sprint(r)
}

func (f Fib3) calc(n int) uint64 {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	result := matrix

	for i := 2; i < n; i++ {
		result = multiply(result, matrix)
	}

	return result[0][0]
}

func multiply(matrixA [2][2]uint64, matrixB [2][2]uint64) [2][2]uint64 {
	var result [2][2]uint64

	result[0][0] = matrixA[0][0]*matrixB[0][0] + matrixA[0][1]*matrixB[1][0]
	result[0][1] = matrixA[0][0]*matrixB[0][1] + matrixA[0][1]*matrixB[1][1]
	result[1][0] = matrixA[1][0]*matrixB[0][0] + matrixA[1][1]*matrixB[1][0]
	result[1][1] = matrixA[1][0]*matrixB[0][1] + matrixA[1][1]*matrixB[1][1]

	return result
}
