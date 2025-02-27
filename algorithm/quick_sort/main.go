package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ARRAY_LENGTH = 5

func main() {
	nums := random_nums(ARRAY_LENGTH)
	fmt.Println("start: ", nums)
	s := time.Now()

	quickSort(nums)

	elapsed := time.Since(s)
	fmt.Println("processimg time: ", elapsed)
}

func sort(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	fmt.Println("  pivot: ", pivot)
	fmt.Println("  low, high: ", low, high)

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	fmt.Println("  in sort: ", nums)
	nums[i+1], nums[high] = nums[high], nums[i+1]
	fmt.Println("  in sort: ", nums)
	return i + 1
}

func reQuickSort(nums []int, low, high int) {
	fmt.Println("mid: ", nums)
	if low < high {
		pivotIndex := sort(nums, low, high)

		reQuickSort(nums, low, pivotIndex-1)
		reQuickSort(nums, pivotIndex+1, high)
	}
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	reQuickSort(nums, 0, len(nums)-1)
}

func random_nums(n int) []int {
	output := make([]int, 0)
	for i := 0; i < n; i++ {
		output = append(output, rand.Intn(ARRAY_LENGTH))
	}
	return output
}
