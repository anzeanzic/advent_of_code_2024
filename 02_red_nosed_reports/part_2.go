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
		//fmt.Println("--------------")
		is_safe := is_report_safe(report)

		if !is_safe {
			//fmt.Println(report)
			//fmt.Println(is_safe)
			//fmt.Println("--------------")
		}

		if is_safe {
			safeCnt++
		}
	}

	fmt.Println(safeCnt)
}

func is_report_safe(report []int) bool {
	report_safe := false
	direction := 0
	allOk := false

	direction = get_direction(report[0], report[1])

	i := 0
	for i = 0; i < len(report)-1; i++ {
		diff := (report[i+1] - report[i]) * direction

		if diff > 0 && diff < 4 {
			allOk = true
		} else {
			allOk = false
			break
		}
	}

	if !allOk && i == len(report)-2 {
		allOk = true
	}

	if !allOk {
		temp := make([]int, len(report))
		copy(temp, report)
		temp = remove(temp, i)

		direction = get_direction(temp[0], temp[1])

		for j := 0; j < len(temp)-1; j++ {
			diff := (temp[j+1] - temp[j]) * direction

			if diff > 0 && diff < 4 {
				allOk = true
			} else {
				allOk = false
				break
			}
		}
	}

	if !allOk {
		temp := make([]int, len(report))
		copy(temp, report)
		temp = remove(temp, i+1)

		direction = get_direction(temp[0], temp[1])

		for j := 0; j < len(temp)-1; j++ {
			diff := (temp[j+1] - temp[j]) * direction

			if diff > 0 && diff < 4 {
				allOk = true
			} else {
				allOk = false
				break
			}
		}
	}

	if !allOk && i != 0 {
		temp := make([]int, len(report))
		copy(temp, report)
		temp = remove(temp, i-1)

		direction = get_direction(temp[0], temp[1])

		for j := 0; j < len(temp)-1; j++ {
			diff := (temp[j+1] - temp[j]) * direction

			if diff > 0 && diff < 4 {
				allOk = true
			} else {
				allOk = false
				break
			}
		}
	}

	if allOk {
		report_safe = true
	}

	return report_safe
}

func is_increasing(num1 int, num2 int) bool {
	return num1 < num2
}

func is_decreasing(num1 int, num2 int) bool {
	return num1 > num2
}

func get_direction(num1 int, num2 int) int {
	direction := 0

	if is_increasing(num1, num2) {
		direction = 1
	} else if is_decreasing(num1, num2) {
		direction = -1
	}

	return direction
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
