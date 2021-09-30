package main

import "fmt"

/**
@author Shitanford
@create 2021-09-26 16:35
*/

func main() {
	s := "abba"
	fmt.Printf("%d\n", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var x = make(map[byte]int)
	max := 0
	for i, j:=0, 0 ; i < len(s); i++ {
		if loc, ok := x[s[i]]; ok {
			if j <= loc {
				j = loc+1
			}
		}
		x[s[i]] = i
		if max < i-j+1 {
			max = i-j+1
		}
	}
	return max
}
