package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	var min, max, sum int64
	min = math.MaxInt64
	max = math.MinInt64
	for i := int64(0); i < n; i++ {
		a := nextInt(sc)
		sum += a
		if min > a {
			min = a
		}
		if max < a {
			max = a
		}
	}

	fmt.Printf("%d %d %d\n", min, max, sum)
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
