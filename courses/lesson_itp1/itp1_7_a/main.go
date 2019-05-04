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

	var valsM, valsF, valsR []int64
	var m, f, r int64
	for {
		m = nextInt(sc)
		f = nextInt(sc)
		r = nextInt(sc)
		if m == -1 && f == -1 && r == -1 {
			break
		}
		valsM = append(valsM, m)
		valsF = append(valsF, f)
		valsR = append(valsR, r)
	}

	for i := range valsM {
		if valsM[i] == -1 || valsF[i] == -1 {
			fmt.Println("F")
			continue
		}
		mfSum := valsM[i] + valsF[i]
		if mfSum >= 80 {
			fmt.Println("A")
		} else if mfSum >= 65 {
			fmt.Println("B")
		} else if mfSum >= 50 {
			fmt.Println("C")
		} else if mfSum < 30 {
			fmt.Println("F")
		} else if valsR[i] >= 50 {
			fmt.Println("C")
		} else {
			fmt.Println("D")
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
