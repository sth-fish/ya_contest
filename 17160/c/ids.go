package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	ids, err := findMissingIDs(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	delim := ""
	for _, id := range ids {
		fmt.Fprintf(os.Stdout, "%s%d", delim, id)
		delim = " "
	}
}

func findMissingIDs(input *bufio.Reader) ([]int, error) {
	buf, err := readFullLine(input)
	if err != nil {
		return nil, fmt.Errorf("can't read from stdin %v", err)
	}

	n, err := strconv.Atoi(string(buf))
	if err != nil {
		return nil, fmt.Errorf("can't parse number %v", err)
	}

	buf, err = readFullLine(input)
	if err != nil {
		return nil, fmt.Errorf("can't read from stdin %v", err)
	}

	ids := make([]int, n)
	for _, val := range strings.Split(string(buf), " ") {
		id, err := strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("can't parse number %v", err)
		}
		ids[id-1] = id
	}

	missingIDs := make([]int, 0, 32)
	for i := 0; i < n; i++ {
		if ids[i] == 0 {
			missingIDs = append(missingIDs, i+1)
		}
	}

	return missingIDs, nil
}

func readFullLine(reader *bufio.Reader) ([]byte, error) {
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

	return buf.Bytes(), err
}
