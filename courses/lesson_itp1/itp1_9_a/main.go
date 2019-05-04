package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	w := nextString(sc)
	var cnt int
	for {
		t := nextString(sc)
		if t == "END_OF_TEXT" {
			break
		}
		if strings.ToLower(t) == w {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
