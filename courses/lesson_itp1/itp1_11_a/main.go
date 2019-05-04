package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type dice struct {
	t int64
	e int64
	w int64
	s int64
	n int64
	b int64
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	d := &dice{
		t: nextInt(sc),
		s: nextInt(sc),
		e: nextInt(sc),
		w: nextInt(sc),
		n: nextInt(sc),
		b: nextInt(sc),
	}

	d.roll(nextString(sc))
	d.Print()
}

func (d *dice) Print() {
	fmt.Println(d.t)
}

func (d *dice) roll(s string) {
	for _, v := range s {
		switch v {
		case 'S':
			d.rollS()
		case 'W':
			d.rollW()
		case 'E':
			d.rollE()
		case 'N':
			d.rollN()
		}
	}
}

func (d *dice) rollS() {
	oldB := d.b
	d.b = d.s
	d.s = d.t
	d.t = d.n
	d.n = oldB
}

func (d *dice) rollN() {
	oldB := d.b
	d.b = d.n
	d.n = d.t
	d.t = d.s
	d.s = oldB
}

func (d *dice) rollE() {
	oldB := d.b
	d.b = d.e
	d.e = d.t
	d.t = d.w
	d.w = oldB
}

func (d *dice) rollW() {
	oldB := d.b
	d.b = d.w
	d.w = d.t
	d.t = d.e
	d.e = oldB
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
