package main

import (
	"bufio"
	"fmt"
	"github.com/mdn010/adventofcode/common"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	common.Checkerror(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//identify highest x and y, can be from separate points
	var maxx, maxy int
	minx, miny := 10000, 10000
	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		common.Checkerror(err)
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		if x < minx {
			minx = x
		}
		if y < miny {
			miny = y
		}
	}

	var mapping = map[string][][]int{}
	for i, line := range lines {
		unit := string(int(`A`[0]) + i)
		mapping[unit] = make([][]int, maxy)
		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		common.Checkerror(err)
		for row := miny; row < maxy; row++ {
			mapping[unit][row] = make([]int, maxx)
			for col := minx; col < maxx; col++ {
				(mapping[unit])[row][col] = int(math.Abs(float64(row-y)) + math.Abs(float64((col - x))))
			}
		}
	}

	finalmap := make([][]string, maxy)

	infinites := make(map[string]int)
	//get infinites
	for i, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		common.Checkerror(err)
		unit := string(int(`A`[0]) + i)

		if x == maxx || y == maxy || x == minx || y == miny {
			infinites[unit] = 1
		}
	}

	for row := miny; row < maxy; row++ {
		finalmap[row] = make([]string, maxx)
		for col := minx; col < maxx; col++ {
			var minunit string
			minval := maxx + maxy
			temp := make(map[int][]string)
			for i, _ := range lines {
				unit := string(int(`A`[0]) + i)
				if minval >= (mapping[unit])[row][col] {
					minval = (mapping[unit])[row][col]
					minunit = unit
					temp[minval] = append(temp[minval], minunit)
				}
			}
			if len(temp[minval]) < 2 {
				finalmap[row][col] = minunit
			}
		}
	}

	finalcount := make(map[string]int)

	for row, v1 := range finalmap {
		for col, _ := range v1 {
			unit := finalmap[row][col]
			if len(unit) != 0 {
				finalcount[unit] += 1
			}
		}
	}

	var finalfinalcount int
	var answer string
	for k, v := range finalcount {
		if _, ok := infinites[k]; !ok {
			if finalfinalcount < v {
				finalfinalcount = v
				answer = k
			}
		}
	}
	//fmt.Println(finalmap)
	fmt.Println(finalfinalcount)
	fmt.Println(answer)
}
