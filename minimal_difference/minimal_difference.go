package main

import "math"

func main() {
	A := []int{3, 1, 2, 4, 3}
	println(solution(A))
}

func solution(A []int) int {
	leftSum := float64(A[0])
	rightSum := float64(0)
	for i := 1; i < len(A); i++ {
		rightSum += float64(A[i])
	}
	minDiff := math.Abs(leftSum - rightSum)
	for i := 1; i < len(A)-1; i++ {
		leftSum += float64(A[i])
		rightSum -= float64(A[i])
		diff := math.Abs(leftSum - rightSum)
		if diff < minDiff {
			minDiff = diff
		}
	}
	return int(minDiff)
}
