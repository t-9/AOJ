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

	s := make([]string, n)
	for x := int64(0); x < n; x++ {
		s[x] = strconv.FormatInt(a[x], 10)
	}
	fmt.Println(strings.Join(s, " "))

	for i := int64(1); i < n; i++ {
		v := a[i]
		j := i - 1
		for j >= 0 && a[j] > v {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = v

		for x := int64(0); x < n; x++ {
			s[x] = strconv.FormatInt(a[x], 10)
		}
		fmt.Println(strings.Join(s, " "))
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
