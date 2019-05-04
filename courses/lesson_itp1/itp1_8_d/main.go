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

	s := nextString(sc)
	p := nextString(sc)

	if strings.Contains(s+s, p) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
