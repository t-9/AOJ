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

	r := nextInt(sc)
	c := nextInt(sc)

	rowSum := make([]int64, r)
	colSum := make([]int64, c)
	var sum int64
	t := make([][]int64, r)

	for i := int64(0); i < r; i++ {
		t[i] = make([]int64, c)
		for j := int64(0); j < c; j++ {
			t[i][j] = nextInt(sc)
			rowSum[i] += t[i][j]
			colSum[j] += t[i][j]
		}
		sum += rowSum[i]
	}

	for i, tr := range t {
		vStr := make([]string, c+1)
		vStr[c] = strconv.FormatInt(rowSum[i], 10)
		for j, v := range tr {
			vStr[j] = strconv.FormatInt(v, 10)
		}
		fmt.Println(strings.Join(vStr, " "))
	}
	vStr := make([]string, c+1)
	vStr[c] = strconv.FormatInt(sum, 10)
	for i, v := range colSum {
		vStr[i] = strconv.FormatInt(v, 10)
	}
	fmt.Println(strings.Join(vStr, " "))
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
