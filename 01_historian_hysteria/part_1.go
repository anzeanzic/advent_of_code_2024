package main

import (
	"bufio"
	"fmt"
	"math"
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

	diffSum := 0

	for len(firstList) >= 1 {
		smallestFirst := 1000000
		smallestSecond := 1000000
		smallestFirstNdx := -1
		smallestSecondNdx := -1

		// find the smallest from first and second array
		for i, _ := range firstList {
			num := firstList[i]

			if num < smallestFirst {
				smallestFirst = num
				smallestFirstNdx = i
			}

			num = secondList[i]

			if num < smallestSecond {
				smallestSecond = num
				smallestSecondNdx = i
			}
		}

		diff := int(math.Abs(float64(smallestFirst - smallestSecond)))
		diffSum += diff

		fmt.Println(smallestFirst, smallestSecond)
		firstList = remove(firstList, smallestFirstNdx)
		secondList = remove(secondList, smallestSecondNdx)

		fmt.Println(diff)
	}

	fmt.Println(diffSum)
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
