package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := random_nums(1000)
	s := time.Now()

	remerge(nums)

	elapsed := time.Since(s)
	fmt.Println("processimg time: ", elapsed)
}

func remerge(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	i := len(nums) / 2
	left := remerge(nums[:i])
	right := remerge(nums[i:])

	return merge(left, right)
}

// left, rightにはソート済みの数字配列を期待する
func merge(left, right []int) []int {
	output := make([]int, 0)
	roops := len(left) + len(right)

	for i := 0; i < roops; i++ {
		if len(left) == 0 {
			output = append(output, right...)
			break
		}
		if len(right) == 0 {
			output = append(output, right...)
			break
		}
		if left[0] > right[0] {
			output = append(output, right[0])
			right = right[1:]
		} else {
			output = append(output, left[0])
			left = left[1:]
		}
	}

	return output
}

func random_nums(n int) []int {
	output := make([]int, 0)
	for i := 0; i < n; i++ {
		output = append(output, rand.Intn(1000))
	}
	return output
}
