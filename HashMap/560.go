package main

import "fmt"

/**
@author Shitanford
@create 2021-09-30 16:29
*/

func main() {
	nums := []int {1,2,1,2,1}
	fmt.Println(subarraySum(nums, 3))
}

//func subarraySum(nums []int, k int) int {
//	ans := 0
//	for i := 0; i < len(nums); i++ {
//		sum := 0
//		for j := i; j < len(nums); j++ {
//			sum += nums[j]
//			if sum == k {
//				ans += 1
//			}
//		}
//	}
//	return ans
//}

func subarraySum(nums []int, k int) int {
	ans := 0
	sum := 0
	sumNum := make(map[int]int)
	sumNum[0] = 1
	for _, value := range nums {
		sum += value
		if num, exits := sumNum[sum-k]; exits {
			ans += num
		}
		if _, exists := sumNum[sum]; exists {
			sumNum[sum] += 1
		} else {
			sumNum[sum] = 1
		}
	}
	return ans
}
