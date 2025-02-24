package main

import "fmt"

func main() {
	nums := []int{2, 5, 3, 8, 7, 1}
	dnums := make([][]int, 0)
	// output := make([]int, len(nums))
	// numa, numb := split(nums)
	// fmt.Println(output, nums[1:])

	for _, v := range nums {
		vv := []int{v}
		dnums = append(dnums, vv)
	}
	fmt.Println(dnums)
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
