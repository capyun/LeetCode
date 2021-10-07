package main

/**
@author Shitanford
@create 2021-10-02 15:34
*/

func main() {

}

func numTrees(n int) int {
	arr := make([]int, n+1)
	arr[0] = 1
	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			arr[i] += arr[k-1] * arr[i-k]
		}
	}
	return arr[n]
}