package main

import (
	"testing"
)

func TestSolve(t *testing.T) {

	var a float64 = 2
	var b float64 = 3
	var c float64 = 4

	var expected1 float64 = 1
	var expected2 float64 = 2

	solution1, solution2 := solve(a, b, c)

    if (solution1 != expected1) {
		t.Errorf("ERROR %s %s", solution1, expected1)
	}

	if (solution2 != expected2) {
		t.Errorf("ERROR %s %s", solution2, expected2)
	}

}
