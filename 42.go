package main

import "fmt"

/**
@author Shitanford
@create 2021-04-18 14:58
*/

func main() {
	arr := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Printf("%d", trap(arr))
}

//方法一、类似动态规划
//leftHighest数组存放在计算下标i时，左边块的最大高度，。。。
//然后逐个计算
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var sum = 0
	var i int
	leftHighest := make([]int, len(height))
	rightHighest := make([]int, len(height))
	leftHighest[1] = height[0]
	rightHighest[len(height)-2] = height[len(height)-1]
	for i=2; i<len(height)-1; i++ {
		if leftHighest[i-1] < height[i-1] {
			leftHighest[i] = height[i-1]
		} else {
			leftHighest[i] = leftHighest[i-1]
		}
	}
	for i=len(height)-3; i>0; i-- {
		if rightHighest[i+1] < height[i+1] {
			rightHighest[i] = height[i+1]
		} else {
			rightHighest[i] = rightHighest[i+1]
		}
	}
	for i=1; i<len(height)-1; i++ {
		if leftHighest[i]>height[i] && rightHighest[i]>height[i] {
			if leftHighest[i] < rightHighest[i] {
				sum += leftHighest[i] - height[i]
			} else {
				sum += rightHighest[i] - height[i]
			}
		}
	}
	return sum
}

//方法二、双指针法
//leftHighest是从左往右的，所以从左往右的下标left可以知道leftHighest, 同理，从右往左的下标right可以知道rightHighest，
//为什么height[left]<height[right]，则leftHighest<rightHighest，
//因为highest的值是由height[left]和height[right]改变的，归纳证明：
//对于初始情况，lh=height[0]，th=height[len(height)-1]，
//如果height[0]<height[len(height)-1]，则计算left，右边至少有height[len(height)-1]比我大
//考虑height[left]<height[right]，lh[left-1]<rh[right]，所以lh<rh
func trap2(height []int) int {
	if len(height) < 3 {
		return 0
	}
	var sum = 0
	left := 0
	right := len(height)-1
	lh, rh := height[0], height[len(height)-1]
	for left < right {
		if height[left] < height[right] {
			sum += lh - height[left]
			left++
			if height[left] > lh {
				lh = height[left]
			}
		} else {
			sum += rh - height[right]
			right--
			if height[right] > rh {
				rh = height[right]
			}
		}
	}
	return sum
}