package main

import (
	"testing"
	"fmt"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1,2,3, 3, 3, 2}
	unique := removeDuplicates(input)
	fmt.Println(unique)
	if len(unique) != 3 {
		t.Errorf("Result not unique") 
	}
}
