package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ARRAY_LENGTH = 10000

func main() {
	nums := random_nums(ARRAY_LENGTH)
	s := time.Now()

	quick_sort(nums)

	elapsed := time.Since(s)
	fmt.Println("processimg time: ", elapsed)
}

func quick_sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	bi := rand.Intn(len(nums))
	b := nums[bi]

	left, right := make([]int, 0), make([]int, 0)
	for _, v := range nums {
		if b > v {
			left = append(left, v)
		} else if b < v {
			right = append(right, v)
		}
	}
	return append(append(quick_sort(left), b), quick_sort(right)...)
}

func random_nums(n int) []int {
	output := make([]int, 0)
	for i := 0; i < n; i++ {
		output = append(output, rand.Intn(ARRAY_LENGTH))
	}
	return output
}
