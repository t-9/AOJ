package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type square struct {
	h int64
	w int64
}

func main() {
	const frame = "#"
	const inside = "."

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var in []*square
	for {
		s := &square{
			nextInt(sc),
			nextInt(sc),
		}

		if s.h == 0 && s.w == 0 {
			break
		}
		in = append(in, s)
	}

	var results []string
	for _, s := range in {
		line := strings.Repeat(frame, int(s.w)) + "\n"
		inLine := frame + strings.Repeat(inside, int(s.w)-2) + frame + "\n"
		set := line + strings.Repeat(inLine, int(s.h)-2) + line
		results = append(results, set)
	}

	fmt.Println(strings.Join(results, "\n"))
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
