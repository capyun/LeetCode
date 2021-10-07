package main

import "fmt"

/**
@author Shitanford
@create 2021-10-02 15:57
*/

func main() {
	arr := [][]int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(arr)
	fmt.Println(arr)
}

func rotate(matrix [][]int)  {
	size := len(matrix)
	for i := 0; i < size/2; i++ {
		matrix[i], matrix[size-i-1] = matrix[size-i-1], matrix[i]
	}
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}