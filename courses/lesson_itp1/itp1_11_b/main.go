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

	q := nextInt(sc)
	for i := int64(0); i < q; i++ {
		d.setTopAndSouth(nextInt(sc), nextInt(sc))
		d.PrintE()
	}
}

func (d *dice) PrintE() {
	fmt.Println(d.e)
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
