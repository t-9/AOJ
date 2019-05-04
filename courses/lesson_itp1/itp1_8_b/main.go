package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var vals []string
	for {
		v := nextString(sc)
		if v == "0" {
			break
		}
		vals = append(vals, v)
	}

	for _, v := range vals {
		var sum int64
		for _, w := range v {
			sum += int64(w - '0')
		}
		fmt.Println(sum)
	}
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
