package main

import (
	"bufio"
	"fmt"
	"github.com/mdn010/adventofcode/common"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	common.Checkerror(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[string]int)

	for scanner.Scan() {
		input := scanner.Text()

		original := input
		count := len(input)

		for x := 0; x < 26; x++ {
			input = original
			unit := string(int(`A`[0]) + x)
			input = strings.Replace(input, unit, "", -1)
			input = strings.Replace(input, strings.ToLower(unit), "", -1)
			for i := 0; i < count; i++ {
				if i+1 < len(input) {
					i = check(&input, i)
				}
			}
			m[unit] = len(input)
		}

		min := len(original)
		var minkey string
		for k, v := range m {
			if v < min {
				min = v
				minkey = k
			}
		}
		fmt.Printf("min: %d, char: %s\n", min, minkey)
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
