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
	m := nextInt(sc)

	mat := make([][]int64, n)
	for i := int64(0); i < n; i++ {
		mat[i] = make([]int64, m)
		for j := int64(0); j < m; j++ {
			mat[i][j] = nextInt(sc)
		}
	}

	vct := make([]int64, m)
	for i := int64(0); i < m; i++ {
		vct[i] = nextInt(sc)
	}

	for i := int64(0); i < n; i++ {
		var x int64
		for j := int64(0); j < m; j++ {
			x += mat[i][j] * vct[j]
		}
		fmt.Println(x)
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
