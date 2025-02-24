package main

import "fmt"

func main() {
	// A := []int{1, 3, 6, 4, 1, 2}
	A := []int{1, 2, 3}
	// A := []int{-1, -3}
	fmt.Println(solution(A))
}

func solution(A []int) int {
	// Create a map to keep track of the elements in the array
	elements := make(map[int]bool)
	// Iterate over the array and add the elements to the map
	for i := 0; i < len(A); i++ {
		elements[A[i]] = true
	}
	// Iterate over the positive integers starting from 1
	for i := 1; i < len(A)+2; i++ {
		// If the integer is not in the map, return it
		if _, ok := elements[i]; !ok {
			return i
		}
	}
	// If all positive integers are present in the array, return 1
	fmt.Println("got here")
	return 1
}