package DFS

/**
@author Shitanford
@create 2021-10-02 20:09
*/

func main() {

}

var ans []string

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	helper(n, 0, "")
	return ans
}

//n is number of ( need to add
//m is number of ( is not matched
func helper(n int, m int, brackets string)  {
	if m == 0 && n == 0 {
		ans = append(ans, brackets)
		return
	}
	if n != 0 {
		helper(n-1, m+1, brackets+"(")
	}
	if m != 0 {
		helper(n, m-1, brackets+")")
	}
}