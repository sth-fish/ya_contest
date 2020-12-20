package main

import (
	"testing"
)

func TestSpiralTraverse(t *testing.T) {
	{
		n := 3
		matrix := [][]int16{
			{9, 2, 3},
			{8, 1, 4},
			{7, 6, 5},
		}

		actual := spiralTraverse(n, matrix)
		expected := []int16{1, 2, 3, 4, 5, 6, 7, 8, 9}

		if len(actual) != len(expected) {
			t.Errorf("got len(actual) %d, want %v", len(actual), len(expected))
		}

		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("got actual[%d]=%d, want expected[%d]=%d", i, actual[i], i, expected[i])
			}
		}
	}

	{
		n := 5
		matrix := [][]int16{
			{4, 10, 7, 10, 9},
			{5, 9, 0, 9, 8},
			{8, 3, 6, 0, 2},
			{8, 10, 3, 0, 0},
			{0, 9, 0, 7, 4},
		}

		actual := spiralTraverse(n, matrix)
		expected := []int16{6, 0, 9, 0, 0, 3, 10, 3, 9, 10, 7, 10, 9, 8, 2, 0, 4, 7, 0, 9, 0, 8, 8, 5, 4}

		if len(actual) != len(expected) {
			t.Errorf("got len(actual) %d, want %v", len(actual), len(expected))
		}

		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("got actual[%d]=%d, want expected[%d]=%d", i, actual[i], i, expected[i])
			}
		}
	}

	{
		n := 1
		matrix := [][]int16{
			{4},
		}

		actual := spiralTraverse(n, matrix)
		expected := []int16{4}

		if len(actual) != len(expected) {
			t.Errorf("got len(actual) %d, want %v", len(actual), len(expected))
		}

		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("got actual[%d]=%d, want expected[%d]=%d", i, actual[i], i, expected[i])
			}
		}
	}
}
