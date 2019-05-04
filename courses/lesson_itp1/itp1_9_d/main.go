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

	s := nextString(sc)
	n := nextInt(sc)
	for i := int64(0); i < n; i++ {
		o := nextString(sc)
		switch o {
		case "print":
			fmt.Println(s[nextInt(sc) : nextInt(sc)+1])
		case "reverse":
			p1 := nextInt(sc)
			p2 := nextInt(sc)
			for p1 < p2 {
				c1 := s[p1]
				t := s[:p1] + string(s[p2]) + s[p1+1:p2] + string(c1)
				if p2+1 <= int64(len(s))-1 {
					t += s[p2+1:]
				}
				s = t
				p1++
				p2--
			}
		case "replace":
			p1 := nextInt(sc)
			p2 := nextInt(sc)
			c := nextString(sc)
			t := s[:p1] + c
			if p2+1 <= int64(len(s))-1 {
				t += s[p2+1:]
			}
			s = t
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
