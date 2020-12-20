// spiral matrix traversal
// n = 3
// ↑ → →
// ↑ ↑ ↓
// ↑ ← ←

// n = 5
// ↑ → → → →
// ↑ ↑ → → ↓
// ↑ ↑ ↑ ↓ ↓
// ↑ ↑ ← ← ↓
// ↑ ← ← ← ←

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	buf, err := readFullLine(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("can't read from stdin %v", err))
		os.Exit(1)
	}

	n, err := strconv.Atoi(string(buf.Bytes()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("can't parse number %v", err))
		os.Exit(1)
	}

	matrix := make([][]int16, n)

	for i := range matrix {
		matrix[i] = make([]int16, n)

		buf, err := readFullLine(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("can't read from stdin %v", err))
			os.Exit(1)
		}

		for j, data := range bytes.Split(buf.Bytes(), []byte(" ")) {
			n, err := strconv.ParseInt(string(data), 10, 16)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", fmt.Errorf("can't parse number %v", err))
				os.Exit(1)
			}

			matrix[i][j] = int16(n)
		}
	}

	output := &bytes.Buffer{}
	output.Grow(n * n)

	for _, val := range spiralTraverse(n, matrix) {
		fmt.Fprintf(output, "%d\n", val)
	}

	fmt.Fprint(os.Stdout, output)
}

func spiralTraverse(n int, matrix [][]int16) []int16 {
	r := make([]int16, 0, n*n)

	center := int(math.Ceil(float64(n)/2)) - 1
	x, y := center-1, center+1

	r = append(r, matrix[center][center])

	for i := 0; i < n/2; i++ {
		for j := x + 1; j <= y; j++ {
			r = append(r, matrix[x][j])
		}

		for j := x + 1; j <= y; j++ {
			r = append(r, matrix[j][y])
		}

		for j := y - 1; j >= x; j-- {
			r = append(r, matrix[y][j])
		}

		for j := y - 1; j >= x; j-- {
			r = append(r, matrix[j][x])
		}

		x--
		y++
	}

	return r
}

func readFullLine(reader *bufio.Reader) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	var (
		line     []byte
		isPrefix bool  = true
		err      error = nil
	)

	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		buf.Write(line)
	}

	return buf, err
}
