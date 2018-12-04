package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]int)

	content, err := ioutil.ReadFile("input.txt")
	//content, err := ioutil.ReadFile("test2.txt")
	checkError(err)

	lines := strings.Split(string(content), "\n")

	for _, v := range lines {
		if len(v) == 0 {
			break
		}
		//fmt.Printf("%v", v)
		care := v[strings.IndexByte(v, '@')+2:]
		//fmt.Println(care)
		key := care[:strings.IndexByte(care, ':')]
		fmt.Println(key)
		xoff := key[:strings.IndexByte(key, ',')]
		//fmt.Println(xoff)
		yoff := key[strings.IndexByte(key, ',')+1:]
		//fmt.Println(yoff)
		area := care[strings.IndexByte(care, ':')+2:]
		//fmt.Println(area)
		w := area[:strings.IndexByte(area, 'x')]
		//fmt.Println(w)
		h := area[strings.IndexByte(area, 'x')+1:]
		//fmt.Println(h)

		//m[key] += 1
		xoffset, err := strconv.Atoi(xoff)
		checkError(err)
		wtoi, err := strconv.Atoi(w)
		checkError(err)

		for x := 1; x <= wtoi; x++ {
			xkey := strconv.Itoa(xoffset + x)

			yoffset, err := strconv.Atoi(yoff)
			checkError(err)
			htoi, err := strconv.Atoi(h)
			checkError(err)
			for y := 1; y <= htoi; y++ {
				ykey := strconv.Itoa(yoffset + y)
				lakey := xkey + "," + ykey
				//fmt.Printf("lakey: %v\n", lakey)
				m[lakey] += 1
			}
		}
	}

	var counter int
	for k, v := range m {
		//fmt.Printf("result: %v %v\n", k, v)
		if v >= 2 {
			counter += 1
			fmt.Printf("result: %v %v\n", k, v)
		}
	}

	fmt.Println(counter)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
