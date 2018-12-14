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

	baseline = "...................." + baseline + "...................................................." //ya like whatever man
	temp := baseline
	mark := 20
	for x := 0; x < 2002; x++ {
		padpots(&baseline, &mark)
		temp = baseline
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
			realcount += i - mark
		}
	}
	fmt.Println(temp)
	fmt.Println(realcount)

	//delta for me is 91 once stabalized
	//2000 generations sum = 184111
	fmt.Println(184111 + (50000000000-2000)*91)
}

func padpots(baseline *string, mark *int) {
	if string([]rune(*baseline)[len(*baseline)-1]) == `#` || string([]rune(*baseline)[len(*baseline)-2]) == `#` || string([]rune(*baseline)[len(*baseline)-3]) == `#` || string([]rune(*baseline)[len(*baseline)-4]) == `#` || string([]rune(*baseline)[len(*baseline)-1]) == `#` {
		*baseline = *baseline + "............."
	}
}
