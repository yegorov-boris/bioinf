package main

import (
	"bufio"
	"fmt"
	"os"
)

type Service struct {
	seq            []byte
	matrix         [][]int
	hairpinMinSize int
}

func main() {
	var seq []byte

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		seq = scanner.Bytes()
	}

	n := len(seq)
	hairpinMinSize := 3

	service := Service{
		seq:            seq,
		matrix:         make([][]int, n, n),
		hairpinMinSize: hairpinMinSize,
	}

	for i := range service.matrix {
		service.matrix[i] = make([]int, n, n)
	}

	service.fill()
	fmt.Print(service.matrix[0][n-1])

	//for i := range service.matrix {
	//	fmt.Printf("%c %v\n", service.seq[i], service.matrix[i])
	//}

	//service.traceBack(0, n-1)
}

func isPair(a, b byte) bool {
	if a < b {
		return a == 'A' && b == 'U' || a == 'C' && b == 'G'
	}

	return b == 'A' && a == 'U' || b == 'C' && a == 'G'
}

func (s *Service) traceBack(i, j int) {
	if j <= i {
		return
	}

	score := s.matrix[i][j]

	if score == s.matrix[i+1][j] {
		s.traceBack(i+1, j)
		return
	}

	if score == s.matrix[i][j-1] {
		s.traceBack(i, j-1)
		return
	}

	if isPair(s.seq[i], s.seq[j]) && j-i > s.hairpinMinSize && score == s.matrix[i+1][j-1]+1 {
		fmt.Println(i+1, j+1) //start numbering from 1
		s.traceBack(i+1, j-1)
		return
	}

	for k := i+1; k < j-1; k++ {
		if score == s.matrix[i][k] + s.matrix[k+1][j] {
			s.traceBack(i, k)
			s.traceBack(k+1, j)
			return
		}
	}
}

func (s *Service) fill() {
	n := len(s.matrix)
	for t := 1; t < n; t++ {
		for i := 0; i < n - t; i++ {
			j := t + i
			max := s.matrix[i+1][j]
			left := s.matrix[i][j-1]
			diag := s.matrix[i+1][j-1]

			if left > max {
				max = left
			}

			if isPair(s.seq[i], s.seq[j]) && j-i > s.hairpinMinSize && diag+1 > max {
				max = diag + 1
			}

			for k := i+1; k < j-1; k++ {
				m := s.matrix[i][k] + s.matrix[k+1][j]
				if m > max {
					max = m
				}
			}

			s.matrix[i][j] = max
		}
	}
}

//GGGAAAUCC
//AAACAUGAGGAUUACCCAUGU
//ACCAAGGGUUGGAAC
