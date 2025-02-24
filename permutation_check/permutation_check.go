package main

import "fmt"

func main() {
	A := []int{4, 1, 3, 2}
	fmt.Println(solution(A))
}

func solution(A []int) int {
	// Create a map to keep track of the elements in the array
	elements := make(map[int]bool)
	// Iterate over the array and add the elements to the map
	for i := 0; i < len(A); i++ {
		elements[A[i]] = true
	}
	// Iterate from 1 to check for missing elements
	for i := 1; i < len(A); i++ {
		// If the integer is not in the map, return it
		if _, ok := elements[i]; !ok {
			return 0
		}
	}
	// If all positive integers are present in the array, return 1
	return 1
}