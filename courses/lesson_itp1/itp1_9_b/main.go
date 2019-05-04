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

	for {
		s := nextString(sc)
		if s == "-" {
			break
		}
		m := nextInt(sc)
		for i := int64(0); i < m; i++ {
			h := nextInt(sc)
			s = s[h:] + s[:h]
		}
		fmt.Println(s)
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
