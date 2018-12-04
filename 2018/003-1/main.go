package main

import (
	"adventofcode/common"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]int)

	content, err := ioutil.ReadFile("input.txt")
	checkError(err)

	lines := strings.Split(string(content), "\n")

	for _, v := range lines {
		if len(v) == 0 {
			break
		}

		var id, xoff, yoff, w, h int
		_, err := fmt.Sscanf(v, "#%d @ %d,%d: %dx%d", &id, &xoff, &yoff, &w, &h)
		checkError(err)

		for x := 1; x <= w; x++ {
			xkey := strconv.Itoa(xoff + x)

			for y := 1; y <= h; y++ {
				ykey := strconv.Itoa(yoff + y)
				lakey := xkey + "," + ykey
				m[lakey] += 1
			}
		}

	}

	var counter int
	for _, v := range m {
		if v >= 2 {
			counter += 1
		}
	}

	fmt.Println(counter)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
