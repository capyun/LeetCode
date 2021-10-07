package main

import "fmt"

/**
@author Shitanford
@create 2021-10-03 10:51
*/

func main() {
	arr := []int {1, 2, 3, 4, 5}
	fmt.Println(subsets(arr))
}

//func subsets(nums []int) [][]int {
//	ans := make([][]int, 1)
//	helper(nums, []int{}, &ans)
//	return ans
//}
//
//func helper(nums []int, arr []int, ans *[][]int)  {
//	if len(nums) == 0{
//		return
//	}
//	for i, num := range nums {
//		x := make([]int, len(arr)+1)
//		copy(x, append(arr, num))
//		fmt.Println(x)
//		*ans = append(*ans, x)
//		helper(nums[i+1: ], x, ans)
//	}
//}

//func subsets(nums []int) [][]int {
//	ans := make([][]int, 1)
//	for _, v := range nums {
//		for _, arr := range ans {
//			x := make([]int, len(arr)+1)
//			copy(x, append(arr, v))
//			ans = append(ans, arr)
//		}
//	}
//	return ans
//}

func subsets(nums []int) [][]int {
	ans := make([][]int, 1)
	bitmask := bits.
	for _, v := range nums {
		for _, arr := range ans {
			x := make([]int, len(arr)+1)
			copy(x, append(arr, v))
			ans = append(ans, arr)
		}
	}
	return ans
}