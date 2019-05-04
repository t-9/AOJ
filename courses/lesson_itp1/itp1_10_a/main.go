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

	x1 := nextFloat(sc)
	y1 := nextFloat(sc)
	x2 := nextFloat(sc)
	y2 := nextFloat(sc)

	xDis := x2 - x1
	yDis := y2 - y1
	fmt.Println(math.Sqrt(xDis*xDis + yDis*yDis))
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
