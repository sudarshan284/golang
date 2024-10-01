package main

import "fmt"

func Add(x, y float64) float64 {
	return x + y
}

func Subtract(x, y float64) float64 {
	return x - y
}

func Divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	fmt.Println("cannot divide by zero")
	return 0
}

func Multiply(x, y float64) float64 {
	return x * y
}

func main() {
	var x float64
	var y float64
	var symbol string
	fmt.Println("--------This is a simple calculator--------------/n")
	fmt.Print("Enter the first number : ")
	fmt.Scan(&x)
	fmt.Print("Enter the operation from +,-,*,/")
	fmt.Scan(&symbol)
	fmt.Print("Enter the second number....")
	fmt.Scan(&y)
	switch symbol {
	case "+":
		fmt.Println("Result after addition: ", Add(x, y))
	case "-":
		fmt.Println("Result after addition: ", Subtract(x, y))
	case "*":
		fmt.Println("Result after addition: ", Multiply(x, y))
	case "/":
		fmt.Println("Result after addition: ", Divide(x, y))
	default:
		fmt.Println("Invalid operator")
	}
}
