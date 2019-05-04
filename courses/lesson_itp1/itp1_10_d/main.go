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
	x := make([]float64, n)
	y := make([]float64, n)
	for i := int64(0); i < n; i++ {
		x[i] = nextFloat(sc)
	}
	for i := int64(0); i < n; i++ {
		y[i] = nextFloat(sc)
	}

	fmt.Println(minkowski(x, y, 1.0))
	fmt.Println(minkowski(x, y, 2.0))
	fmt.Println(minkowski(x, y, 3.0))
	fmt.Println(chebyshev(x, y))
}

func minkowski(x []float64, y []float64, p float64) float64 {
	var s float64
	for i := 0; i < len(x); i++ {
		s += math.Pow(math.Abs(x[i]-y[i]), p)
	}
	return math.Pow(s, 1.0/p)
}

func chebyshev(x []float64, y []float64) float64 {
	var m float64
	for i := 0; i < len(x); i++ {
		c := math.Abs(x[i] - y[i])
		if c > m {
			m = c
		}
	}
	return m
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
