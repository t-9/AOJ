package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	s := sequenceGets()

	m := make(map[rune]int64, 26)
	for r := rune('a'); r <= 'z'; r++ {
		m[r] = 0
	}

	for _, v := range s {
		for _, v2 := range v {
			w := unicode.ToLower(v2)
			if w < 'a' || w > 'z' {
				continue
			}
			m[w] = m[w] + 1
		}
	}

	for r := rune('a'); r <= 'z'; r++ {
		fmt.Printf("%s : %d\n", string(r), m[r])
	}
}

func sequenceGets() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			break
		}
		lines = append(lines, str)
	}

	return
}
