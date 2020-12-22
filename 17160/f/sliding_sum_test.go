package main

import (
	"testing"
)

func TestSpiralTraverse(t *testing.T) {
	{
		sum := adder(3)

		if ans := sum(1); ans != 1 {
			t.Errorf("got %d, want 1", ans)
		}

		if ans := sum(2); ans != 3 {
			t.Errorf("got %d, want 3", ans)
		}

		if ans := sum(3); ans != 6 {
			t.Errorf("got %d, want 6", ans)
		}

		if ans := sum(4); ans != 9 {
			t.Errorf("got %d, want 9", ans)
		}

		if ans := sum(5); ans != 12 {
			t.Errorf("got %d, want 12", ans)
		}

		if ans := sum(1); ans != 10 {
			t.Errorf("got %d, want 10", ans)
		}
	}
}
