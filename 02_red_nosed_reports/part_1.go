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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Fields(line)
		report := make([]int, 0, len(line_arr))

		for _, numStr := range line_arr {
			num, err := strconv.Atoi(numStr)

			if err != nil {
				panic(err)
			}

			report = append(report, num)
		}

		reports = append(reports, report)
	}

	file.Close()
	fmt.Println("---------------------------------------")

	safeCnt := 0

	for _, report := range reports {
		fmt.Println(report)
		is_safe := is_report_safe(report)
		fmt.Println(is_safe)

		if is_safe {
			safeCnt++
		}
	}

	fmt.Println(safeCnt)
}

func is_report_safe(report []int) bool {
	report_safe := true
	previous := -1
	increasing := false
	decreasing := false

	for i, num := range report {
		if i == 1 {
			if is_increasing(previous, num) {
				increasing = true
			} else if is_decreasing(previous, num) {
				decreasing = true
			} else {
				report_safe = false
				break
			}
		}
		if i != 0 {
			diff := math.Abs(float64(previous - num))

			if diff < 1 || diff > 3 {
				report_safe = false
				break
			}

			if increasing && !is_increasing(previous, num) {
				report_safe = false
				break
			} else if decreasing && !is_decreasing(previous, num) {
				report_safe = false
				break
			}
		}

		previous = num
	}

	return report_safe
}

func is_increasing(num1 int, num2 int) bool {
	return num1 < num2
}

func is_decreasing(num1 int, num2 int) bool {
	return num1 > num2
}
