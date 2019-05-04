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

	n := nextInt(sc)

	var ap, bp int64
	for i := int64(0); i < n; i++ {
		a := nextString(sc)
		b := nextString(sc)
		if a > b {
			ap += 3
		} else if a < b {
			bp += 3
		} else {
			ap++
			bp++
		}
	}
	fmt.Println(ap, bp)
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
