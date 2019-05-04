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

	n := nextInt(sc)

	d := make([]*dice, n)
	for i := int64(0); i < n; i++ {
		d[i] = &dice{
			t: nextInt(sc),
			s: nextInt(sc),
			e: nextInt(sc),
			w: nextInt(sc),
			n: nextInt(sc),
			b: nextInt(sc),
		}
	}

	for i := int64(0); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if d[i].same(d[j]) {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

func (d *dice) same(d2 *dice) bool {
	d2.setTopAndSouth(d.t, d.s)
	return d.t == d2.t && d.b == d2.b && d.n == d2.n &&
		d.e == d2.e && d.s == d2.s && d.w == d2.w
}

func (d *dice) setTopAndSouth(t, s int64) {
	d.setTop(t)
	switch s {
	case d.e:
		d.pitchL()
	case d.n:
		d.pitchR()
		d.pitchR()
	case d.w:
		d.pitchR()
	}
}

func (d *dice) setTop(t int64) {
	switch t {
	case d.s:
		d.rollN()
	case d.n:
		d.rollS()
	case d.e:
		d.rollW()
	case d.w:
		d.rollE()
	case d.b:
		d.rollN()
		d.rollN()
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

func (d *dice) pitchR() {
	oldN := d.n
	d.n = d.e
	d.e = d.s
	d.s = d.w
	d.w = oldN
}

func (d *dice) pitchL() {
	oldN := d.n
	d.n = d.w
	d.w = d.s
	d.s = d.e
	d.e = oldN
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
