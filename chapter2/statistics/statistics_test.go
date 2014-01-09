package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	var x = []float64{1, 2, 3}

	var expected float64 = 6
	var actual float64 = sum(x)

	if (actual != expected) {
		t.Errorf("ERROR %s %s", actual, expected)
	}
}

func TestMedian(t *testing.T) {

	var x = []float64{1, 2, 3}

	var expected float64 = 2
	var actual float64 = median(x)

	if (actual != expected) {
		t.Errorf("ERROR %s %s", actual, expected)
	}
}

func TestMean(t *testing.T) {

	var x = []float64{1, 2, 3}

	var expected float64 = 2
	var actual float64 = mean(x)

	if (actual != expected) {
		t.Errorf("ERROR %s %s", actual, expected)
	}
}

func TestMode(t *testing.T) {

	var x = []float64{1, 2, 2, 3, 3, 5}

	var expected = []float64{2, 3}
	var actual []float64 = mode(x)

	for i := 0; i < len(expected); i++ {
        if (actual[i] != expected[i]) {
			t.Errorf("ERROR %s %s", actual, expected)
		}
    }

}
