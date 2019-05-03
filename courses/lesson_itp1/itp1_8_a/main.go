package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	rd := bufio.NewReaderSize(os.Stdin, 1000000)

	s := readLine(rd)

	var r []rune
	for _, v := range s {
		if unicode.IsUpper(v) {
			r = append(r, unicode.ToLower(v))
		} else {
			r = append(r, unicode.ToUpper(v))
		}
	}
	fmt.Println(string(r))
}

func readLine(rd *bufio.Reader) string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rd.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
