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
	m := make([]int64, n)
	var cnt int64

	for i := int64(0); i < n; i++ {
		m[i] = nextInt(sc)
	}

	for _, v := range m {
		if v == 2 {
			cnt++
			continue
		}
		if v%2 == 0 {
			continue
		}
		sqrtV := math.Sqrt(float64(v))
		for i := int64(3); float64(i) <= sqrtV; i += 2 {
			if v%i == 0 {
				cnt--
				break
			}
		}
		cnt++
	}
	fmt.Println(cnt)
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
