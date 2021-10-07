package Trie

import "strings"

/**
@author Shitanford
@create 2021-10-04 10:20
*/

func main() {

}


func wordBreak(s string, wordDict []string) bool {
	can := make([]int, len(s))
	for i := 0; i < len(can); i++ {
		can[i] = -1
	}
	return helper(s, 0, wordDict, can)
}

func helper(s string, index int, wordDict []string, can []int) bool {
	if len(s) == 0 {
		return true
	}
	if can[index] == 1 {
		return true
	} else if can[index] == 0 {
		return false
	}
	for _, str := range wordDict {
		if strings.HasPrefix(s, str) {
			if helper(s[len(str): ], index+len(str), wordDict, can) {
				can[index] = 1
				return true
			}
		}
	}
	can[index] = 0
	return false
}