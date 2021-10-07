package main

/**
@author Shitanford
@create 2021-09-30 19:41
*/

func main() {

}

//func groupAnagrams(strs []string) [][]string {
//	anagrams := make(map[string][]string)
//	for _, str := range strs {
//		chars := strings.Split(str, "")
//		sort.Strings(chars)
//		key := strings.Join(chars, "")
//		anagrams[key] = append(anagrams[key], str)
//	}
//	ans := make([][]string, 0)
//	for _, strGroup := range anagrams {
//		ans = append(ans, strGroup)
//	}
//	return ans
//}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)
	for _, str := range strs {
		key := charsCount(str)
		anagrams[key] = append(anagrams[key], str)
	}
	var ans [][]string
	for _, strGroup := range anagrams {
		ans = append(ans, strGroup)
	}
	return ans
}

func charsCount(str string) [26]int {
	var count [26]int
	for _, char := range str {
		count[char-'a'] += 1
	}
	return count
}