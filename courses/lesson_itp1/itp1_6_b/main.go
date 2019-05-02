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

	valsNumber := make([]int64, n)
	valsMark := make([]string, n)
	for i := int64(0); i < n; i++ {
		valsMark[i] = nextString(sc)
		valsNumber[i] = nextInt(sc)
	}

	m := make(map[int64]bool, n)

	for i := int64(0); i < n; i++ {
		p := valsNumber[i]
		switch valsMark[i] {
		case "H":
			p += 13
		case "C":
			p += 26
		case "D":
			p += 39
		}
		m[p] = true
	}

	for i := int64(1); i <= 52; i++ {
		if _, ok := m[i]; !ok {
			cn := (i-1)%13 + 1
			var cm string
			if i > 39 {
				cm = "D"
			} else if i > 26 {
				cm = "C"
			} else if i > 13 {
				cm = "H"
			} else {
				cm = "S"
			}
			fmt.Printf("%s %d\n", cm, cn)
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
