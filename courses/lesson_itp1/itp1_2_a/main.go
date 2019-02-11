package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := nextInt(sc)
	b := nextInt(sc)
	fmt.Println(resolve(a, b))
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextInt(sc *bufio.Scanner) int64 {
	i, err := strconv.ParseInt(nextString(sc), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func resolve(a, b int64) string {
	if a == b {
		return "a == b"
	} else if a < b {
		return "a < b"
	} else {
		return "a > b"
	}
}
