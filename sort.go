package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("numbers.txt")

	strContent := string(content)
	strArray := strings.Fields(strContent)
	var strNumArray = []float64{}

	for _, i := range strArray {
		j, _ := strconv.ParseFloat(i, 64)
		strNumArray = append(strNumArray, j)
	}

	//bubble sort algo
	isSorted := false

	for !isSorted {
		isSorted = true
		var i = 0
		for i < len(strNumArray)-1 {
			if strNumArray[i] > strNumArray[i+1] {
				strNumArray[i], strNumArray[i+1] = strNumArray[i+1], strNumArray[i]
				isSorted = false
			}
			i++
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	var sortedStrArray = []string{}
	for _, i := range strNumArray {
		sortedStr := fmt.Sprintf("%g", i)
		sortedStrArray = append(sortedStrArray, sortedStr)
	}

	fmt.Println(sortedStrArray)

	//write to new file
	file, err := os.OpenFile("sortedNumbers.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	fileWriter := bufio.NewWriter(file)

	for _, data := range sortedStrArray {
		_, _ = fileWriter.WriteString(data + "\n")
	}

	fileWriter.Flush()
	file.Close()

}
