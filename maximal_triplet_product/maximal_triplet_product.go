package main

func main() {
	// A := []int{4, 5, 1, 0}
	A := []int{-3, 1, 2, -2, 5, 6}
	println(solution(A))
}

func solution(A []int) int {
	maximalProduct := 0
	arrangedArray := make([]int, len(A))
	copy(arrangedArray, A)
	// Pick only consequtive ascending numbers
	for i := 1; i < len(arrangedArray); i++ {
		if arrangedArray[i] <= arrangedArray[i-1] {
			arrangedArray = append(arrangedArray[:i], arrangedArray[i+1:]...)
			i-=1
		}
	}
	for i := 0; i < len(arrangedArray)-2; i++ {
		maximalProduct = max(maximalProduct, arrangedArray[i]*arrangedArray[i+1]*arrangedArray[i+2])
	}

	return maximalProduct
}
