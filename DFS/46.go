package main

import "fmt"

/**
@author Shitanford
@create 2021-10-02 22:11
*/

func main() {
	permute([]int {1, 2, 3})
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	helper(nums, 0, &ans)
	return ans
}

func helper(nums []int, index int, ans *[][]int)  {
	if index == len(nums) {
		arr := make([]int, index)
		copy(arr, nums)
		fmt.Println(arr)
		*ans = append(*ans, arr)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		helper(nums, index+1, ans)
		nums[i], nums[index] = nums[index], nums[i]
	}
}