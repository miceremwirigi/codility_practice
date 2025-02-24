package main

import (
	"fmt"
	"slices"
)

func main() {
	// A := []int{10, 2, 5, 1, 8, 20}
	
	A := []int{10, 50, 5, 1}
	for i := 0; i < len(A); i++ {
		print(A[i], " ")
	}
	println("")
	println(solution(A))
}

func solution(A []int) int {
	hasTriangularTriplet := 0
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
	// fmt.Println(arrangedArray)

	// Check if there is a triangular triplet
	// for i := 0; i < len(arrangedArray)-2; i++ {
	// 	if arrangedArray[i]+arrangedArray[i+1] > arrangedArray[i+2] &&
	// 		arrangedArray[i+1]+arrangedArray[i+2] > arrangedArray[i] &&
	// 		arrangedArray[i+2]+arrangedArray[i] > arrangedArray[i+1] {
	// 		hasTriangularTriplet = 1
	// 		fmt.Println(arrangedArray[i], arrangedArray[i+1], arrangedArray[i+2])
	// 		break
	// 	}
	// }
	for i := 0; i < len(B)-2; i++ {
		if B[i]+B[i+1] > B[i+2] &&
			B[i+1]+B[i+2] > B[i] &&
			B[i+2]+B[i] > B[i+1] {
			hasTriangularTriplet = 1
			fmt.Println(B[i], B[i+1], B[i+2])
			break
		}
	}
	return hasTriangularTriplet
}
