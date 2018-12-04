package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input002.txt")
	checkError(err)

	reader := bufio.NewReaderSize(f, 1024*1024)

	var twos, threes int64

	for {
		id := readLine(reader)
		if id == "" {
			break
		}

		m := make(map[string]int64)

		for _, r := range id {
			c := string(r)
			v := m[c]
			m[c] = v + 1
		}

		var flagtwos, flagthrees bool
		for _, v := range m {
			if v == 2 {
				flagtwos = true
			}
			if v == 3 {
				flagthrees = true
			}
		}

		if flagtwos {
			twos += 1
		}
		if flagthrees {
			threes += 1
		}
		fmt.Printf("loop twos %v\n", twos)
		fmt.Printf("loop twos %v\n", threes)
	}
	fmt.Println(twos)
	fmt.Println(threes)
	fmt.Println(twos * threes)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
