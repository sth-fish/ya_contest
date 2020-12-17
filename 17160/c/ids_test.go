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
		expexted := 2
		actual := len(ids)

		if actual != expexted {
			t.Errorf("got len(ids) %d, want %d", actual, expexted)
		}
	}

	{
		expexted := 5
		actual := ids[0]

		if expexted != actual {
			t.Errorf("got ids[0] = %d, want %d", actual, expexted)
		}
	}

	if ids[0] != 5 {
		t.Errorf("got ids[0] = %d, want 5", ids[0])
	}

	if ids[1] != 7 {
		t.Errorf("got ids[1] = %d, want 7", ids[1])
	}
}
