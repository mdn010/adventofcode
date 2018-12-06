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

	finalmap := make([][]int, maxy)

	var finalcount int

	for row := miny; row < maxy; row++ {
		finalmap[row] = make([]int, maxx)
		for col := minx; col < maxx; col++ {
			var tempdist int
			for i, _ := range lines {
				unit := string(int(`A`[0]) + i)
				tempdist += (mapping[unit])[row][col]
			}
			if 10000 > tempdist {
				finalcount += 1
			}
		}
	}

	fmt.Println(finalcount)
}
