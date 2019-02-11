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
	v := nextInt(sc)
	h := nextInt(sc)

	area := v * h
	perimeter := v + h + v + h

	fmt.Println(area, perimeter)
}

func nextInt(sc *bufio.Scanner) int64 {
	sc.Scan()
	i, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}
