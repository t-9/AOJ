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

	n := nextInt(sc)

	call(n)
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

func call(n int64) {
	checkNum(1, n)
}

func checkNum(i, n int64) (int64, int64) {
	x := i
	if x%3 == 0 {
		fmt.Printf(" %d", i)
		return endCheckNum(x, i, n)
	}
	return include3(x, i, n)
}

func include3(x, i, n int64) (int64, int64) {
	if x%10 == 3 {
		fmt.Printf(" %d", i)
		return endCheckNum(x, i, n)
	}
	x /= 10
	if x != 0 {
		return include3(x, i, n)
	}
	return endCheckNum(x, i, n)
}

func endCheckNum(x, i, n int64) (int64, int64) {
	i++
	if i <= n {
		return checkNum(i, n)
	}
	fmt.Println("")
	return x, i
}
