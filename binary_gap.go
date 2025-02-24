package main

import (
	"fmt"
)

func main() {
	// N := 2147483647
	N := 25781
	fmt.Printf("N: %d %b\n", N, N)

	fmt.Println(solution(N))

}

func solution (N int) int {
    currentGapLength := 0
    maxGapLength := 0
    inGap := false

    for N > 0 {
        // Check if the value, 2, 3, of least significant bit is 1
        if N&1 == 1 {
            if inGap {
                if currentGapLength > maxGapLength {
                    maxGapLength = currentGapLength
                }
                currentGapLength = 0
            } else {
                inGap = true
            }
        } else {
            if inGap {
                currentGapLength++
            }
        }
        // Right shift N by 1 bit
        N >>= 1
    }

    return maxGapLength
}
