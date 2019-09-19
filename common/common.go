package common

import (
	"os"
	"bufio"
	"io"
	"strings"
)

func Checkerror(err error) {
	if err != nil {
		panic(err)
	}
}

func NewFile(filename string) *os.File {
	file, err := os.Open(filename)
	Checkerror(err)
	return file
}

func NewBufferedReader(filename string) *bufio.Reader {
	f := NewFile(filename)
	reader := bufio.NewReaderSize(f, 1024*1024)
	return reader
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
