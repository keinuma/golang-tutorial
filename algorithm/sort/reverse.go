package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	fmt.Println("original: ", src)

	fmt.Println("output of reverseAlpha: ", reverseAlpha(src))
	fmt.Println("output of reverseBeta: ", reverseBeta(src))

	fmt.Println("original: ", src)
}

func reverseAlpha(src []int) []int {
	output := make([]int, len(src))
	copy(output, src)

	for i := 0; i < len(src)/2; i++ {
		output[i], output[len(src)-i-1] = src[len(src)-i-1], src[i]
	}
	return output
}

func reverseBeta(src []int) []int {
	output := make([]int, len(src))
	copy(output, src)

	for left, right := 0, len(src)-1; left < right; left, right = left+1, right-1 {
		output[left], output[right] = src[right], src[left]
	}
	return output
}
