package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator")
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
