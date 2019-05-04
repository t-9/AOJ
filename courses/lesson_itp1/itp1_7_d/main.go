package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	m := nextInt(sc)
	l := nextInt(sc)

	a := make([][]int64, n)
	for i := int64(0); i < n; i++ {
		a[i] = make([]int64, m)
		for j := int64(0); j < m; j++ {
			a[i][j] = nextInt(sc)
		}
	}

	b := make([][]int64, m)
	for i := int64(0); i < m; i++ {
		b[i] = make([]int64, l)
		for j := int64(0); j < l; j++ {
			b[i][j] = nextInt(sc)
		}
	}

	for i := int64(0); i < n; i++ {
		o := make([]string, l)
		for j := int64(0); j < l; j++ {
			var v int64
			for k := int64(0); k < m; k++ {
				v += a[i][k] * b[k][j]
			}
			o[j] = strconv.FormatInt(v, 10)
		}
		fmt.Println(strings.Join(o, " "))
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
