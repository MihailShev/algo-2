package fibonacci

// Расчет чисел Фибоначчи через возведение матрицы в степень

import (
	"fmt"
	"math/big"
	"strconv"
)

var startMatrix = [2][2]*big.Int{[2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(0)}}

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

func (f Fib3) calc(n int) *big.Int {
	if n == 0 {
		return zero
	}

	if n <= 2 {
		return one
	}

	result := startMatrix

	var res [2][2]*big.Int

	pow := n - 1

	for pow > 0 {
		if pow%2 == 1 {
			if res[0][0] == nil {
				res = result
			} else {
				res = multiply(res, result)
			}
		}

		result = multiply(result, result)
		pow /= 2
	}

	return res[0][0]
}

func (f Fib3) calc2(n int) *big.Int {
	if n == 0 {
		return zero
	}

	if n <= 2 {
		return one
	}

	result := startMatrix

	for i := 2; i < n; i++ {
		result = multiply(result, startMatrix)
	}

	return result[0][0]
}

func multiply(matrixA [2][2]*big.Int, matrixB [2][2]*big.Int) [2][2]*big.Int {
	var result [2][2]*big.Int

	tmp := big.NewInt(0).Mul(matrixA[0][0], matrixB[0][0])
	tmp.Add(tmp, big.NewInt(0).Mul(matrixA[0][1], matrixB[1][0]))

	result[0][0] = tmp

	tmp = big.NewInt(0).Mul(matrixA[0][0], matrixB[0][1])
	tmp = tmp.Add(tmp, big.NewInt(0).Mul(matrixA[0][1], matrixB[1][1]))

	result[0][1] = tmp

	tmp = big.NewInt(0).Mul(matrixA[1][0], matrixB[0][0])
	tmp = tmp.Add(tmp, big.NewInt(0).Mul(matrixA[1][1], matrixB[1][0]))

	result[1][0] = tmp

	tmp = big.NewInt(0).Mul(matrixA[0][1], matrixB[0][1])
	tmp = tmp.Add(tmp, big.NewInt(0).Mul(matrixA[1][1], matrixB[1][1]))

	result[1][1] = tmp

	return result
}
