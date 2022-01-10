package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	f, err := os.Open("numbers.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	total := 0.0
	scanner := bufio.NewScanner(f)

	//add up total in for loop
	for scanner.Scan() {
		strNum := scanner.Text()
		actualNum, _ := strconv.ParseFloat(strNum, 64)
		total += actualNum
	}

	//output total value from numbers.txt
	fmt.Println("Sum of numbers.txt is: ", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
