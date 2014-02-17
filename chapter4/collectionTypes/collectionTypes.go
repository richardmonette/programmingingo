package main

func In(haystack []int, needle int) bool {
	for i := range haystack {
		if (haystack[i] == needle) {
			return true
		}
	}
	return false
}

func removeDuplicates(input []int) []int {
	result := []int{}
	for i := range input {
		if (In(result, input[i]) == false) {
			result = append(result, input[i])
		}
	}
	return result
}
