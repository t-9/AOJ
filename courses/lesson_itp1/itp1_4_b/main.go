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
	r := nextFloat(sc)

	x := r * r * math.Pi
	y := r * 2 * math.Pi
	fmt.Printf("%f %f\n", x, y)
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextFloat(sc *bufio.Scanner) float64 {
	f, err := strconv.ParseFloat(nextString(sc), 64)
	if err != nil {
		panic(err)
	}
	return f
}
