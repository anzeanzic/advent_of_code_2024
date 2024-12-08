package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	memory := ""

	for scanner.Scan() {
		memory += scanner.Text()
	}

	r, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := r.FindAllStringIndex(memory, -1)
	sum := 0

	fmt.Println(len(matches))

	for _, match := range matches {
		str_start, str_end := match[0], match[1]
		str := memory[str_start:str_end]
		fmt.Println(str)
		str = strings.Replace(str, "mul(", "", 1)
		str = strings.Replace(str, ")", "", 1)
		num_arr := strings.Split(str, ",")

		num1, err1 := strconv.Atoi(num_arr[0])
		num2, err2 := strconv.Atoi(num_arr[1])

		if err1 != nil && err2 != nil {
			panic(err)
		}

		sum += num1 * num2
	}

	file.Close()
	fmt.Println("---------------------------------------")
	fmt.Println(sum)
}
