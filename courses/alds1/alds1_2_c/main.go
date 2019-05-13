package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	a := make([]string, n)
	as := make([]string, n)
	an := make([]int64, n)
	for i := int64(0); i < n; i++ {
		a[i] = nextString(sc)
		as[i] = a[i][0:1]
		var err error
		an[i], err = strconv.ParseInt(a[i][1:2], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	ban, bas := bubbleSort(n, an, as)
	san, sas := selectionSort(n, an, as)
	b := make([]string, n)
	s := make([]string, n)
	isSame := true
	for i := int64(0); i < n; i++ {
		b[i] = fmt.Sprintf("%s%d", bas[i], ban[i])
		s[i] = fmt.Sprintf("%s%d", sas[i], san[i])
		if b[i] != s[i] {
			isSame = false
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(b), "[]"))
	fmt.Println("Stable")
	fmt.Println(strings.Trim(fmt.Sprint(s), "[]"))
	if isSame {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

func bubbleSort(n int64, sa []int64, ss []string) ([]int64, []string) {
	a := make([]int64, cap(sa))
	copy(a, sa)

	s := make([]string, cap(ss))
	copy(s, ss)

	isIncomp := true

	for isIncomp {
		isIncomp = false
		for j := n - 1; j >= 1; j-- {
			if a[j] < a[j-1] {
				t := a[j-1]
				a[j-1] = a[j]
				a[j] = t

				ts := s[j-1]
				s[j-1] = s[j]
				s[j] = ts

				isIncomp = true
			}
		}
	}
	return a, s
}

func selectionSort(n int64, sa []int64, ss []string) ([]int64, []string) {
	a := make([]int64, cap(sa))
	copy(a, sa)

	s := make([]string, cap(ss))
	copy(s, ss)

	for i := int64(0); i <= n-1; i++ {
		minJ := i
		for j := int64(i); j <= n-1; j++ {
			if a[j] < a[minJ] {
				minJ = j
			}
		}
		t := a[minJ]
		a[minJ] = a[i]
		a[i] = t

		ts := s[minJ]
		s[minJ] = s[i]
		s[i] = ts
	}
	return a, s
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
