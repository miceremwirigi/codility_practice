package main

import (
	"fmt"
	"slices"
)

func main() {
	// A := []int{4, 5, 1, 0}
	A := []int{-3, 1, 2, -2, 5, 6}
	println(solution(A))
}

func solution(A []int) int {
	maximalProduct := 0
	B := make([]int, len(A))
	copy(B, A)
	// Sort the array
	slices.Sort(B)
	fmt.Println(A)
	fmt.Println(B)

	// arrangedArray := make([]int, len(A))
	// copy(arrangedArray, A)
	// // Pick only consequtive ascending numbers
	// for i := 1; i < len(arrangedArray); i++ {
	// 	if arrangedArray[i] <= arrangedArray[i-1] {
	// 		arrangedArray = append(arrangedArray[:i], arrangedArray[i+1:]...)
	// 		i-=1
	// 	}
	// }
	// for i := 0; i < len(arrangedArray)-2; i++ {
	// 	maximalProduct = max(maximalProduct, arrangedArray[i]*arrangedArray[i+1]*arrangedArray[i+2])
	// }

	// Check if there is a triangular triplet
	for i := 0; i < len(B)-2; i++ {
		maximalProduct = max(maximalProduct, B[i]*B[i+1]*B[i+2])
	}

	return maximalProduct
}
