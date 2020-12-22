package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	nBuf, nErr := readFullLine(input)
	if nErr != nil {
		fmt.Fprintf(os.Stderr, "can't read from stdin %v", nErr)
		os.Exit(1)
	}

	n, nErr := strconv.Atoi(nBuf.String())
	if nErr != nil {
		fmt.Fprintf(os.Stderr, "can't parse int %v", nErr)
		os.Exit(1)
	}

	valsBuf, valsErr := readFullLine(input)
	if valsErr != nil {
		fmt.Fprintf(os.Stderr, "can't read from stdin %v", valsErr)
		os.Exit(1)
	}

	res := 0
	sum := adder(n)
	for _, valBuf := range bytes.Split(valsBuf.Bytes(), []byte(" ")) {
		val, valErr := strconv.Atoi(string(valBuf))
		if valErr != nil {
			fmt.Fprintf(os.Stderr, "can't parse int %v", valErr)
			os.Exit(1)
		}

		res = sum(val)
	}

	fmt.Fprint(os.Stdout, res)
}

func adder(n int) func(val int) int {
	i := 0
	vals := make([]int, n)

	return func(val int) int {
		vals[i] = val
		i = (i + 1) % len(vals)

		sum := 0
		for _, v := range vals {
			sum += v
		}

		return sum
	}
}

func readFullLine(r *bufio.Reader) (*bytes.Buffer, error) {
	b := &bytes.Buffer{}

	var (
		line     []byte
		isPrefix bool  = true
		err      error = nil
	)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()

		b.Write(line)
	}

	return b, err
}
