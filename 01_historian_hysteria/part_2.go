package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ReadFile()
}

func ReadFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var firstList []int
	var secondList []int

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Fields(line)
		firstNumber, err := strconv.Atoi(line_arr[0])
		secondNumber, err2 := strconv.Atoi(line_arr[1])
		fmt.Println(firstNumber, secondNumber)

		if err != nil || err2 != nil {
			panic(err)
		}

		firstList = append(firstList, firstNumber)
		secondList = append(secondList, secondNumber)
	}

	file.Close()
	fmt.Println("---------------------------------------")

	simScore := 0

	for _, num1 := range firstList {
		counter := 0

		for _, num2 := range secondList {
			if num1 == num2 {
				counter++
			}
		}

		fmt.Println(num1, counter)
		simScore += num1 * counter
	}

	fmt.Println(simScore)
}
