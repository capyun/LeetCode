package main

import "fmt"

/**
@author Shitanford
@create 2021-09-27 19:58
*/

func main() {
	a := [][]int {
		{1,1,1,0,0},
		{1,1,1,1,1},
		{0,0,0,0,0},
		{1,1,1,1,1},
		{1,0,1,0,1},
	}
	b := [][]int {
		{0,0,0,0,0},
		{1,1,1,1,1},
		{0,1,0,1,0},
		{0,1,0,1,0},
		{1,0,0,0,1},
	}
	fmt.Println(countSubIslands(a, b))
}

var visited [500][500]bool

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	sum := 0
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[0]); j++ {
			if !isVisited(i, j) && grid2[i][j] == 1 {
				if isSubIsland(grid1, grid2, i, j) {
					sum += 1
				}
			}
		}
	}
	return sum
}

func isSubIsland(grid1 [][]int, grid2 [][]int, i int, j int) bool {
	if i<0 || j<0 || i>=len(grid1) || j>=len(grid1[0]) ||
		isVisited(i, j) || grid2[i][j]==0 {
		return true
	}
	visited[i][j] = true
	var ans = grid1[i][j] == 1
	ans = isSubIsland(grid1, grid2, i, j-1) && ans
	ans = isSubIsland(grid1, grid2, i, j+1)	&& ans
	ans = isSubIsland(grid1, grid2, i-1, j)	&& ans
	ans = isSubIsland(grid1, grid2, i+1, j)	&& ans
	return ans
}

func isVisited(i int, j int) bool {
	return visited[i][j]
}
