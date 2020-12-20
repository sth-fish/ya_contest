package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestFindMissingIDs(t *testing.T) {
	input := []byte(`7
6 4 1 2 3
`)

	r := bufio.NewReader(bytes.NewReader(input))
	ids, err := findMissingIDs(r)
	if err != nil {
		t.Errorf("got error %v", err)
	}

	{
		expected := 2
		actual := len(ids)

		if actual != expected {
			t.Errorf("got len(ids) %d, want %d", actual, expected)
		}
	}

	{
		expected := 5
		actual := ids[0]

		if expected != actual {
			t.Errorf("got ids[0] = %d, want %d", actual, expected)
		}
	}

	{
		expected := 7
		actual := ids[1]

		if expected != actual {
			t.Errorf("got ids[0] = %d, want %d", actual, expected)
		}
	}
}
