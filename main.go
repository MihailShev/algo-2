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
	//fmt.Printf("*** Test power by iteration ***\n\n")
	//powByIteration := tester.NewTester(power.Pow1{}, pathToPowerData)
	//powByIteration.RunTest()
	//
	fmt.Printf("\n*** Test power by power of 2 with multiplication\n\n")
	powByPow2Multiplication := tester.NewTester(power.Pow2{}, pathToPowerData)
	powByPow2Multiplication.RunTest()

	//fmt.Printf("\n*** Test power by binary decomposition of the exponent\n\n")
	//powByBinaryDecomposition := tester.NewTester(power.Pow3{}, pathToPowerData)
	//powByBinaryDecomposition.RunTest()

	fibonacciByRecursion := tester.NewTester(fibonacci.Fib1{}, pathToFibonacciData)
	fibonacciByRecursion.RunTestWithCount(8)
}
