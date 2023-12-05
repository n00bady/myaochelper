package myaochelper

import "strings"

// Takes a slice of strings and a key and returns the index
// of the key if it finds it else it return -1
func FindIndex(lines []string, key string) int {
	for i, l := range lines {
		if strings.Contains(l, key) {
			return i
		}
	}

	return -1
}

// Takes a slice of ints and finds the index of the minimum element
func FindMinIndex(s []int) int {
	min := s[0]
	minIdx := 0
	for i, v := range s {
		if v < min {
			min = v
			minIdx = i
		}
	}

	return minIdx
}

// Similar as FindMinIndex but finds the index of the maximum
func FindMaxIndex(s []int) int {
	max := s[0]
	maxIdx := 0
	for i, v := range s {
		if v > max {
			max = v
			maxIdx = i
		}
	}

	return maxIdx
}
