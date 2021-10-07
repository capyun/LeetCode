package DFS

/**
@author Shitanford
@create 2021-10-02 16:53
*/

func main() {

}

var numberToLetter = []string{"", "", "abc", "def", "ghi", "jkl",
	"mno", "pqrs", "tuv", "wxyz"}
var ans []string

func letterCombinations(digits string) []string {
	ans = make([]string, 0)
	if len(digits) == 0 {
		return ans
	}
	addLetters(digits, 0, "")
	return ans
}

func addLetters(digits string, index int, letters string)  {
	if index == len(digits) {
		ans = append(ans, letters)
		return
	}
	str := numberToLetter[digits[index] - '0']
	for _, char := range str {
		addLetters(digits, index + 1, letters + string(char))
	}
}