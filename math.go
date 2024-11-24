package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Add is our function that sums two integers
func Add(x, y int) int {
	return x + y
}

// Subtract subtracts two integers
func Subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println("Welcome to the Math Operations Application!")

	// Interactive input for testing Add and Subtract
	for {
		fmt.Println("Choose an operation:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// Perform addition
			fmt.Print("Enter the first number: ")
			x, err1 := readInput()
			if err1 != nil {
				fmt.Println(err1)
				continue
			}

			fmt.Print("Enter the second number: ")
			y, err2 := readInput()
			if err2 != nil {
				fmt.Println(err2)
				continue
			}

			fmt.Printf("Result of %d + %d = %d\n", x, y, Add(x, y))
		case 2:
			// Perform subtraction
			fmt.Print("Enter the first number: ")
			x, err1 := readInput()
			if err1 != nil {
				fmt.Println(err1)
				continue
			}

			fmt.Print("Enter the second number: ")
			y, err2 := readInput()
			if err2 != nil {
				fmt.Println(err2)
				continue
			}

			fmt.Printf("Result of %d - %d = %d\n", x, y, Subtract(x, y))
		case 3:
			fmt.Println("Exiting the application. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}

// readInput is a helper function to read and validate integer input
func readInput() (int, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, fmt.Errorf("Error reading input: %v", err)
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("Invalid number. Please enter a valid integer.")
	}

	return num, nil
}
