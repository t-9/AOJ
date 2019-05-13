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
	a := make([]int64, n)
	for i := int64(0); i < n; i++ {
		a[i] = nextInt(sc)
	}

	isIncomp := true
	cnt := 0

	for isIncomp {
		isIncomp = false
		for j := n - 1; j >= 1; j-- {
			if a[j] < a[j-1] {
				t := a[j-1]
				a[j-1] = a[j]
				a[j] = t
				isIncomp = true
				cnt++
			}
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(a), "[]"))
	fmt.Println(cnt)
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
