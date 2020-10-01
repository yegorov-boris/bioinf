package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	ab := strings.Split(text, " ")
	a := ab[0]
	b := ab[1]
	n := len(a)
	m := len(b)
	f := make([][]int, n+1, n+1)

	for i := range f {
		f[i] = make([]int, m+1, m+1)
		f[i][0] = -i
	}

	for j := range f[0] {
		f[0][j] = -j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			match := f[i-1][j-1] + s(a[i-1], b[j-1])
			del := f[i-1][j] - 1
			insert := f[i][j-1] - 1

			max := match
			if del > max {
				max = del
			}
			if insert > max {
				max = insert
			}

			f[i][j] = max
		}
	}

	var alignmentA, alignmentB string //TODO: make it faster with bytes buffer

	i:= n
	j := m

	for ; i > 0 || j > 0; {
		score := f[i][j]
		scoreDiag := f[i - 1][j - 1]
		scoreLeft := f[i - 1][j]

		switch {
		case score == scoreDiag + s(a[i-1], b[j-1]):
			alignmentA = alignmentA + string([]byte{a[i-1]})
			alignmentB = alignmentB + string([]byte{b[j-1]})
			i--
			j--
		case score == scoreLeft - 1:
			alignmentA = alignmentA + string([]byte{a[i-1]})
			alignmentB = alignmentB + "_"
			i--
		default:
			alignmentA = alignmentA + "_"
			alignmentB = alignmentB + string([]byte{b[j-1]})
			j--
		}
	}

	for ; i > 0; i-- {
		alignmentA = alignmentA + string([]byte{a[i-1]})
		alignmentB = alignmentB + "_"
	}

	for ; j > 0; j-- {
		alignmentB = alignmentB + string([]byte{b[j-1]})
		alignmentA = alignmentA + "_"
	}

	fmt.Printf("%s %s\n", reverse(alignmentA), reverse(alignmentB))
}

func s(a, b byte) int {
	if a == b {
		return 1
	}

	return -1
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
