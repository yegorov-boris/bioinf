package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	r := csv.NewReader(strings.NewReader(string(dat)))

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	weights := make(map[string]int, len(records[0]))

	for i, k := range records[0] {
		w, e := strconv.ParseInt(records[1][i], 10, 32)
		if e != nil {
			fmt.Println(err)
			return
		}

		weights[k] = int(w)
	}

	fmt.Println("weights:")
	fmt.Println(weights)

	a := os.Args[2]
	b := os.Args[3]

	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//text := scanner.Text()
	//
	//ab := strings.Split(text, " ")
	//if len(ab) == 1 {
	//	ab = append(ab, "")
	//}
	//
	//a := ab[0]
	//b := strings.Join(ab[1:], " ")

	n := len(a)
	m := len(b)

	if n == 0 {
		alignmentA := make([]byte, m, m)

		for i := range b {
			alignmentA[i] = '_'
		}

		fmt.Printf("%s %s\n", string(alignmentA), b)

		return
	}

	if m == 0 {
		alignmentB := make([]byte, n, n)

		for i := range a {
			alignmentB[i] = '_'
		}

		fmt.Printf("%s %s\n", a, string(alignmentB))

		return
	}

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
			match := f[i-1][j-1] + sw(weights, a[i-1], b[j-1])
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

	// nolint: godox
	var alignmentA, alignmentB string //TODO: make it faster with bytes buffer

	i:= n
	j := m

	for ; i > 0 || j > 0; {
		score := f[i][j]

		switch {
		case i > 0 && j > 0 && score == f[i - 1][j - 1] + sw(weights, a[i-1], b[j-1]):
			alignmentA = alignmentA + string([]byte{a[i-1]})
			alignmentB = alignmentB + string([]byte{b[j-1]})
			i--
			j--
		case i > 0 && score == f[i - 1][j] - 1:
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

	fmt.Println("aligned sequences:")
	fmt.Println(reverse(alignmentA))
	fmt.Println(reverse(alignmentB))
	//fmt.Printf("%s %s\n", reverse(alignmentA), reverse(alignmentB))
}

func s(a, b byte) int {
	if a == b {
		return 1
	}

	return -1
}

func sw(weights map[string]int, a, b byte) int {
	w, ok := weights[string([]byte{a, b})]
	if ok {
		return w
	}

	return s(a, b)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
