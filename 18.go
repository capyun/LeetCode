package main

import "fmt"

/**
@author Shitanford
@create 2021-09-27 15:59
*/

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fourSum(nums, 0)
	for _, arr := range ans {
		fmt.Println(arr)
	}
}

var ans [][]int

func fourSum(nums []int, target int) [][]int {
	var arr []int
	for i := 0; i <= len(nums)-4; i++ {
		fourSum2(0, arr, nums, target, i)
	}
	return ans
}

func fourSum2(sum int, selected []int, nums []int, target int, index int) {
	sum += nums[index]
	selected = append(selected, nums[index])
	if len(selected) == 4 {
		if sum == target {
			x := make([]int, 4)
			copy(x, selected)
			ans = append(ans, x)
		}
		return
	} else {
		for i := index+1; i < len(nums); i++ {
			fourSum2(sum, selected, nums, target, i)
		}
	}
}