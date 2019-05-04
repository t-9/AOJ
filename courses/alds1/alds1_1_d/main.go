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
	r := make([]int64, n)

	for i := int64(0); i < n; i++ {
		r[i] = nextInt(sc)
	}

	var max, min int64
	max = math.MinInt64
	min = r[0]

	for i := int64(1); i < n; i++ {
		t := r[i] - min
		if max < t {
			max = t
		}
		if min > r[i] {
			min = r[i]
		}
	}
	fmt.Println(max)
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
