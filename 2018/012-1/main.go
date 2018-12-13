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

	reader := bufio.NewReader(file)

	baseline := (common.Readline(reader))[15:]

	common.Readline(reader)
	line := common.Readline(reader)

	code := make(map[string]string)
	for line != "" {
		code[line[:5]] = string(line[9])
		line = common.Readline(reader)
	}

	baseline = "...................." + baseline + "...................................." //ya like whatever man
	temp := baseline
	for x := 0; x < 20; x++ {
		for i, _ := range baseline {
			if i < 2 || i > (len(baseline)-3) {
				continue
			}
			if val, ok := code[baseline[i-2:i+3]]; ok {
				temp = temp[:i] + val + temp[i+1:]
			} else {
				temp = temp[:i] + `.` + temp[i+1:]
			}
		}
		baseline = temp
	}

	var realcount int
	for i, v := range baseline {
		if string(v) == `#` {
			realcount += i - 20
		}
	}
	fmt.Println(temp)
	fmt.Println(realcount)
}
