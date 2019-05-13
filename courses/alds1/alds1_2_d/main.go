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

	m, g, cnt := shellSort(a, n)
	fmt.Println(m)
	fmt.Println(strings.Trim(fmt.Sprint(g), "[]"))
	fmt.Println(cnt)
	for i := int64(0); i < n; i++ {
		fmt.Println(a[i])
	}
}

func shellSort(a []int64, n int64) (int64, []int64, int64) {
	var cnt int64
	m, g := calcM(n)
	for i := int64(0); i < m; i++ {
		cnt = insertionSort(a, n, g[i], cnt)
	}
	return m, g, cnt
}

func insertionSort(a []int64, n int64, g int64, cnt int64) int64 {
	for i := g; i <= n-1; i++ {
		v := a[i]
		j := i - g
		for j >= 0 && a[j] > v {
			a[j+g] = a[j]
			j = j - g
			cnt++
		}
		a[j+g] = v
	}
	return cnt
}

func calcM(n int64) (int64, []int64) {
	if n == 1 {
		return 1, []int64{0}
	}
	var cnt int64
	c := int64(1)
	var g []int64
	for c < n {
		g = append(g, c)
		c = c*3 + 1
		cnt++
	}
	for i := len(g)/2 - 1; i >= 0; i-- {
		opp := len(g) - 1 - i
		g[i], g[opp] = g[opp], g[i]
	}
	return cnt, g
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
