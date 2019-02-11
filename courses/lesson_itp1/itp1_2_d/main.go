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
	w := nextInt(sc)
	h := nextInt(sc)
	x := nextInt(sc)
	y := nextInt(sc)
	r := nextInt(sc)
	fmt.Println(resolve(w, h, x, y, r))
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

func resolve(w, h, x, y, r int64) string {
	const no = "No"
	const yes = "Yes"
	if x-r < 0 {
		return no
	}
	if x+r > w {
		return no
	}
	if y-r < 0 {
		return no
	}
	if y+r > h {
		return no
	}
	return yes
}
