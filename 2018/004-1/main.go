package main

import (
	"bufio"
	"fmt"
	"github.com/mdn010/adventofcode/common"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	var currentelf int
	currentnaptime := time.Now()
	var naptimetotal = map[int]int{}
	var naptimes = map[int]map[int]int{}
	//m := make(map[int]make(map[int]int)) //elf -> minute -> counter
	file, err := os.Open("input.txt")
	common.Checkerror(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.SliceStable(lines, func(i, j int) bool {
		li := lines[i][1:strings.IndexByte(lines[i], ']')]
		lj := lines[j][1:strings.IndexByte(lines[j], ']')]
		//1518-02-04 00:14
		ti, err := time.Parse("2006-01-02 15:04", li)
		common.Checkerror(err)
		tj, err := time.Parse("2006-01-02 15:04", lj)
		common.Checkerror(err)

		return tj.Sub(ti) > 0
	})

	for _, line := range lines {
		var year, month, day, hour, minute, id int
		_, err := fmt.Sscanf(line[:strings.IndexByte(line, ']')+1], "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)
		common.Checkerror(err)
		//[1518-03-19 00:02] Guard #647 begins shift
		num, err := fmt.Sscanf(line[strings.IndexByte(line, ']')+2:], "Guard #%d begins shift", &id)
		if err == nil && num == 1 {
			currentelf = id
		} else {
			n := line[1:strings.IndexByte(line, ']')]
			t, err := time.Parse("2006-01-02 15:04", n)
			common.Checkerror(err)
			message := line[strings.IndexByte(line, ']')+2:]
			if message == "wakes up" {
				naptime := int(t.Sub(currentnaptime).Minutes())
				naptimetotal[currentelf] += naptime
				for i := 1; i <= naptime; i++ {
					ladate := t.Add(time.Minute * time.Duration(-i))
					if naptimes[currentelf] == nil {
						naptimes[currentelf] = map[int]int{}
					}
					(naptimes[currentelf])[int(ladate.Minute())] += 1
				}
			} else {
				currentnaptime = t
			}
		}
	}

	var sleepyelf, maxslept int
	for elfkey, val := range naptimetotal {
		if val > maxslept {
			sleepyelf = elfkey
			maxslept = val
		}
	}

	fmt.Printf("elf: %d\n", sleepyelf)

	sleepyminute := 61
	var sleepval int
	for minute, val := range naptimes[sleepyelf] {
		if val >= sleepval {
			sleepyminute = minute
			sleepval = val
		} else if val == sleepval && minute < sleepyminute {
			sleepyminute = minute
			sleepval = val
		}
	}
	fmt.Printf("minute: %d\n", sleepyminute)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
