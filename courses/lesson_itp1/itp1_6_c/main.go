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

	valsB := make([]int64, n)
	valsF := make([]int64, n)
	valsR := make([]int64, n)
	valsV := make([]int64, n)
	for i := int64(0); i < n; i++ {
		valsB[i] = nextInt(sc)
		valsF[i] = nextInt(sc)
		valsR[i] = nextInt(sc)
		valsV[i] = nextInt(sc)
	}

	m := make(map[int64]int64)
	for i := int64(0); i < n; i++ {
		k := valsB[i]*1000 + valsF[i]*100 + valsR[i]
		v, ok := m[k]
		if !ok {
			v = 0
		}
		m[k] = v + valsV[i]
	}

	var out string
	for b := int64(1); b <= 4; b++ {
		if b != 1 {
			out += "####################\n"
		}
		for f := int64(1); f <= 3; f++ {
			for r := int64(1); r <= 10; r++ {
				out += " "

				k := b*1000 + f*100 + r
				v, ok := m[k]
				if !ok {
					v = 0
				}
				out += strconv.FormatInt(v, 10)
				if r == 10 {
					out += "\n"
				}
			}
		}
	}
	fmt.Print(out)
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
