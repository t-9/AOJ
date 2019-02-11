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

const (
	white = '#'
	black = '.'
)

func main() {

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

	var boards [][]string
	for _, s := range in {
		boards = append(boards, s.getBoard())
	}

	var results []string
	for _, b := range boards {
		results = append(results, strings.Join(b, "\n")+"\n")
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

func (sq *square) getBoard() []string {
	l := make([]string, sq.h)
	for y := int64(0); y < sq.h; y++ {
		l[y] = getLine(y, sq.w)
	}
	return l
}

func getLine(y, w int64) string {
	r := make([]rune, w)
	for x := int64(0); x < w; x++ {
		r[x] = getDot(x, y)
	}
	return string(r)
}

func getDot(x, y int64) rune {
	if (x+y)%2 == 0 {
		return white
	}
	return black
}
