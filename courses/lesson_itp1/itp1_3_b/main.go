package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var in []string
	for {
		v := nextString(sc)
		if v == "0" {
			break
		}
		in = append(in, v)
	}

	for idx, v := range in {
		fmt.Printf("Case %d: %v\n", idx+1, v)
	}
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
