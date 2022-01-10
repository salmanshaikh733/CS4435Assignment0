package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//make scanner object
	reader := bufio.NewScanner(os.Stdin)

	for {
		//instructions for user
		fmt.Println("Enter a math operation followed by two number, enter 'stop' to stop the program")
		//get user input
		reader.Scan()
		//store as slice of array
		input := strings.Fields(reader.Text())

		//to stop the program
		if input[0] == "stop" {
			os.Exit(1)
		}

		//parse to a floating point
		num1, _ := strconv.ParseFloat(input[1], 64)
		num2, _ := strconv.ParseFloat(input[2], 64)

		//simple math operations
		if input[0] == "add" {
			sum := num1 + num2
			fmt.Println("The sum is:")
			fmt.Println(sum)
		} else if input[0] == "sub" {
			dif := num1 - num2
			fmt.Println("The difference is: ")
			fmt.Println(dif)

		} else if input[0] == "div" {
			//if division by zero return error, not possible
			if num2 == 0 {
				fmt.Println("Error divsion by zero")
			} else {
				div := num1 / num2
				fmt.Println("The division is:")
				fmt.Println(div)
			}

		} else if input[0] == "mult" {
			product := num1 * num2
			fmt.Println("The product is: ")
			fmt.Println(product)

		} else {
			fmt.Println("Invalid input, try again...")
		}
	}
}
