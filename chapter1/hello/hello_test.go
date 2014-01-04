package main

import "testing"

func TestPrinter(t *testing.T) {

	const expected = "Hello World!"

	actual := printer()

	if (actual != expected) {
		t.Errorf("ERROR %s %s", actual, expected)
	}

}
