package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := 110201
	//input := 2018
	recipes := "37"
	elf1 := 0
	elf2 := 1

	for i := 0; i < input+10; i++ {
		recipe1, _ := strconv.Atoi(string(recipes[elf1]))
		recipe2, _ := strconv.Atoi(string(recipes[elf2]))

		sum := strconv.Itoa(recipe1 + recipe2)

		recipes += sum
		elf1 += int(math.Mod(float64(recipe1+1), float64(len(recipes))))
		elf2 += int(math.Mod(float64(recipe2+1), float64(len(recipes))))
		if elf1 >= len(recipes) {
			elf1 = int(math.Mod(float64(elf1), float64(len(recipes))))
		}
		if elf2 >= len(recipes) {
			elf2 = int(math.Mod(float64(elf2), float64(len(recipes))))
		}
	}

	fmt.Println(recipes[input:(input + 10)])
}
