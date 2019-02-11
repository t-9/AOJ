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
	var in [][]int64
	for {
		x := nextInt(sc)
		y := nextInt(sc)
		if x == 0 && y == 0 {
			break
		}
		in = append(in, []int64{x, y})
	}

	for _, l := range in {
		var min, max int64
		if l[0] > l[1] {
			min = l[1]
			max = l[0]
		} else {
			min = l[0]
			max = l[1]
		}
		fmt.Printf("%d %d\n", min, max)
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
