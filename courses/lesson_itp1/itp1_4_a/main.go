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
	a := nextFloat(sc)
	b := nextFloat(sc)

	x := int64(a) / int64(b)
	y := int64(a) % int64(b)
	z := a / b
	fmt.Printf("%d %d %f\n", x, y, z)
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
