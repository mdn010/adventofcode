package main

import (
	"bufio"
	"fmt"
	"github.com/mdn010/adventofcode/common"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	common.Checkerror(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()

		count := len(input)
		for i := 0; i < count; i++ {
			if i+1 < len(input) {
				i = check(&input, i)
			}
		}

		fmt.Println(len(input))
	}
}

func check(input *string, i int) int {
	if ((*input)[i]+32) == (*input)[i+1] || ((*input)[i]-32) == (*input)[i+1] {
		*input = (*input)[:i] + (*input)[i+2:]
		if i == 0 {
			return -1
		} else {
			return i - 2
		}
	}
	return i
}
