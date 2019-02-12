package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	vals := make([]int64, n)
	for i := int64(0); i < n; i++ {
		vals[i] = nextInt(sc)
	}

	strVals := make([]string, n)
	var cnt int64
	for i := n - 1; i >= 0; i-- {
		strVals[i] = strconv.FormatInt(vals[cnt], 10)
		cnt++
	}

	fmt.Println(strings.Join(strVals, " "))
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
