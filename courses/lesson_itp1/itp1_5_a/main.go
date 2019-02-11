package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const grid = "#"

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var in [][]int64
	for {
		h := nextInt(sc)
		w := nextInt(sc)

		if h == 0 && w == 0 {
			break
		}
		in = append(in, []int64{h, w})
	}

	var results []string
	for _, l := range in {
		line := strings.Repeat(grid, int(l[1])) + "\n"
		set := strings.Repeat(line, int(l[0]))
		results = append(results, set)
	}

	fmt.Println(strings.Join(results, "\n"))
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
