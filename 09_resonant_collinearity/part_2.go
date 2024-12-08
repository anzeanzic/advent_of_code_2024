package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var antinode_antennas_overlap = make(map[string]int)

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
	var antennas [][]string

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, "")
		antennaLine := make([]string, 0, len(line_arr))

		for _, char := range line_arr {
			antennaLine = append(antennaLine, char)
		}

		antennas = append(antennas, antennaLine)
	}

	file.Close()
	PrintAntennas(antennas)
	fmt.Println()
	CreateAntinodes(antennas)
}

func CreateAntinodes(antennas [][]string) {
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[i]); j++ {
			if antennas[i][j] != "." && antennas[i][j] != "#" {
				//fmt.Println(antennas[i][j], i, j)
				//fmt.Println("------------")
				FindNextAntenna(antennas, i, j, antennas[i][j])
				//PrintAntennas(antennas)
			}
		}
	}

	fmt.Println("------------")
	PrintAntennas(antennas)
	fmt.Println(antinode_antennas_overlap)
	fmt.Println(len(antinode_antennas_overlap))
}

func FindNextAntenna(antennas [][]string, i_start int, j_start int, antenna_sign string) {
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[i]); j++ {
			if antennas[i][j] == antenna_sign {
				new_x, new_y := NewPoint(i_start, j_start, i, j)

				if new_x >= 0 && new_x < len(antennas[i]) && new_y >= 0 && new_y < len(antennas) {
					antinode_antennas_overlap[strconv.Itoa(new_x)+","+strconv.Itoa(new_y)] = antinode_antennas_overlap[strconv.Itoa(new_x)+strconv.Itoa(new_y)] + 1

					if antennas[new_x][new_y] == "." {
						antennas[new_x][new_y] = "#"
					}
				}

				p1_x := i
				p1_y := j
				p2_x := new_x
				p2_y := new_y

				for f := 0; f < 100; f++ {
					new_x, new_y := NewPoint(p1_x, p1_y, p2_x, p2_y)

					if new_x >= 0 && new_x < len(antennas[i]) && new_y >= 0 && new_y < len(antennas) {
						antinode_antennas_overlap[strconv.Itoa(new_x)+","+strconv.Itoa(new_y)] = antinode_antennas_overlap[strconv.Itoa(new_x)+strconv.Itoa(new_y)] + 1

						if antennas[new_x][new_y] == "." {
							antennas[new_x][new_y] = "#"
						}

						p1_x = p2_x
						p1_y = p2_y
						p2_x = new_x
						p2_y = new_y
					} else {
						break
					}
				}
			}
		}
	}
}

func NewPoint(x1 int, y1 int, x2 int, y2 int) (int, int) {
	x3 := x1 + 2*(x2-x1)
	y3 := y1 + 2*(y2-y1)

	return x3, y3
}

func PrintAntennas(antennas [][]string) {
	for i := 0; i < len(antennas); i++ {
		for j := 0; j < len(antennas[i]); j++ {
			fmt.Print(antennas[i][j])
		}

		fmt.Println()
	}
}
