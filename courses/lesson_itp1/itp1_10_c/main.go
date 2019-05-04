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

	for {
		n := nextInt(sc)
		if n == 0 {
			break
		}
		s := make([]float64, n)
		var sum float64
		for i := int64(0); i < n; i++ {
			s[i] = nextFloat(sc)
			sum += s[i]
		}
		m := sum / float64(n)

		var t float64
		for i := int64(0); i < n; i++ {
			t += math.Pow(s[i]-m, 2.0)
		}
		fmt.Println(math.Sqrt(t / float64(n)))
	}
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

func nextFloat(sc *bufio.Scanner) float64 {
	i, err := strconv.ParseFloat(nextString(sc), 64)
	if err != nil {
		panic(err)
	}
	return i
}
