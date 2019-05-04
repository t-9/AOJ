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

	var valsN, valsX []int64
	var n, x int64
	for {
		n = nextInt(sc)
		x = nextInt(sc)
		if n == 0 && x == 0 {
			break
		}
		valsN = append(valsN, n)
		valsX = append(valsX, x)
	}

	m := make([]map[int64]int64, 98)
	for a := int64(3); a <= 100; a++ {
		m[a-3] = make(map[int64]int64)
		if a-4 >= 0 {
			for i := int64(0); i <= 300; i++ {
				cnt, ok := m[a-4][i]
				if ok {
					m[a-3][i] = cnt
				} else {
					m[a-3][i] = 0
				}
			}
		}
		for b := int64(1); b <= a-2; b++ {
			for c := b + 1; c < a; c++ {
				d := a + b + c
				cnt, ok := m[a-3][d]
				if ok {
					m[a-3][d] = cnt + 1
				} else {
					m[a-3][d] = 1
				}
			}
		}
	}

	for i := range valsN {
		v, ok := m[valsN[i]-3][valsX[i]]
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println(0)
		}
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
