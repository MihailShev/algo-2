package main

import (
	"fmt"
	"github.com/MihailShev/algo-2/v1/fibonacci"
	"github.com/MihailShev/algo-2/v1/power"
	tester "github.com/MihailShev/algo-tester"
)

const pathToPowerData = "test-data/power"
const pathToFibonacciData = "test-data/fibonacci"

func main() {
	fmt.Printf("*** Test power by iteration ***\n\n")
	powByIteration := tester.NewTester(power.Pow1{}, pathToPowerData)
	powByIteration.RunTestWithCount(10)

	fmt.Printf("\n*** Test power by power of 2 with multiplication ***\n\n")
	powByPow2Multiplication := tester.NewTester(power.Pow2{}, pathToPowerData)
	powByPow2Multiplication.RunTestWithCount(10)

	fmt.Printf("\n*** Test power by binary decomposition of the exponent ***\n\n")
	powByBinaryDecomposition := tester.NewTester(power.Pow3{}, pathToPowerData)
	powByBinaryDecomposition.RunTestWithCount(10)

	fmt.Printf("\n*** Fibonacci by recursion ***\n\n")
	fibonacciByRecursion := tester.NewTester(fibonacci.Fib1{}, pathToFibonacciData)
	fibonacciByRecursion.RunTestWithCount(10)

	fmt.Printf("\n*** Fibonacci by golden ratio ***\n\n")
	fibonacciByGoldenRation := tester.NewTester(fibonacci.Fib2{}, pathToFibonacciData)
	fibonacciByGoldenRation.RunTestWithCount(13)

	fmt.Printf("\n*** Fibonacci by matrix ***\n\n")
	fibonacciByMatrix := tester.NewTester(fibonacci.Fib3{}, pathToFibonacciData)
	fibonacciByMatrix.RunTestWithCount(13)

	fmt.Printf("\n*** Finish ***\n")
	_, _ = fmt.Scanf(" ")
}
