package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input002.txt")
	checkError(err)

	lines := strings.Split(string(content), "\n")
	//fmt.Println(len(lines))
	//os.Exit(3)

	for x_i, x := range lines {
		for y_i, y := range lines {
			if len(y) == 0 {
				break
			}
			var diffs int
			for i, c := range x {
				fmt.Println(x)
				fmt.Println(y)
				fmt.Println(i)
				if string(c) != string(y[i]) {
					diffs += 1
				}
			}
			if diffs == 1 {
				fmt.Println(lines[x_i])
				fmt.Println(lines[y_i])
				os.Exit(3)
			}
		}
	}

	/*var twos, threes int64

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
	*/
}

/*func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}*/

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
