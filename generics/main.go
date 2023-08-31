package main

import "fmt"

// Without generics
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// With generics
func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func main() {
	firstInts, secondInts := splitAnySlice([]string{"abc", "def", "ghi", "pqr"})
	fmt.Println(firstInts, secondInts)
}
