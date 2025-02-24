package main

func main(){
	A := []int{9, 3, 9, 3, 9, 7, 9}
	println(solution(A))
}

func solution(A []int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		result ^= A[i] // find XOR of all elements
	}
	return result
}