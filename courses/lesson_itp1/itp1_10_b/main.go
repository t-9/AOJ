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

	a := nextFloat(sc)
	b := nextFloat(sc)
	c := nextFloat(sc)

	cRad := c * math.Pi / 180

	h := b * math.Sin(cRad)
	s := a * h / 2.0
	l := a + b + math.Sqrt(a*a+b*b-2*a*b*math.Cos(cRad))

	fmt.Println(s)
	fmt.Println(l)
	fmt.Println(h)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextFloat(sc *bufio.Scanner) float64 {
	i, err := strconv.ParseFloat(nextString(sc), 64)
	if err != nil {
		panic(err)
	}
	return i
}
