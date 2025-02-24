package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := random_nums(100)
	dnums := make([][]int, 0)
	for _, v := range nums {
		vv := []int{v}
		dnums = append(dnums, vv)
	}

	s := time.Now()
	output := remerge(dnums)
	elapsed := time.Since(s)

	fmt.Println("input: ", nums)
	fmt.Println("output: ", output)
	fmt.Println("processimg time: ", elapsed)
}

func remerge(nums [][]int) []int {
	if len(nums) == 1 {
		return nums[0]
	}

	mnums := make([][]int, 0)
	for i := 0; i < len(nums); i = i + 2 {
		if len(nums)-1 == i {
			mnums = append(mnums, nums[len(nums)-1])
			break
		}
		numa, numb := nums[i], nums[i+1]
		mnums = append(mnums, merge(numa, numb))
	}

	return remerge(mnums)
}

// numa, numaにはソート済みの数字配列を期待する
func merge(numa, numb []int) []int {
	output := make([]int, 0)
	roops := len(numa) + len(numb)

	for i := 0; i < roops; i++ {
		if len(numa) == 0 {
			output = append(output, numb...)
			break
		}
		if len(numb) == 0 {
			output = append(output, numa...)
			break
		}
		if numa[0] > numb[0] {
			output = append(output, numb[0])
			numb = numb[1:]
		} else {
			output = append(output, numa[0])
			numa = numa[1:]
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
