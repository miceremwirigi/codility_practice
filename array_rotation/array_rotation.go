package main

func main(){
	A := []int{3, 8, 9, 7, 6}
	K := 3
	for i := 0; i < len(A); i++ {
		print(solution(A, K)[i], " ")
	}
	println()
	println(solution(A, K))
}

func solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	
	K = K % len(A)
	
	if K == 0 {
		return A
	}
	
	result := make([]int, len(A))
	
	for i := 0; i < len(A); i++ {
		result[(i + K) % len(A)] = A[i]
	}
	
	return result
}