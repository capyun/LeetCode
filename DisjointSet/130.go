package DisjointSet

/**
@author Shitanford
@create 2021-10-06 21:53
*/

func main() {

}

type islandSet map[int]int

func solve(board [][]byte)  {
	set := make(islandSet)
	cols := len(board[0])
	for i := 0; i < len(board); i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				set.add(i, j, cols)
			}
		}
	}
}

func coordinate(i, n int) (int, int) {
	return i / n, i % n
}

func index(x, y, n int) int {
	return x * n + y
}

func (this islandSet) add(x, y, n int)  {
	i := index(x, y, n)
	if x==0
}

func (this islandSet) initial(board [][]byte)  {
	cols := len(board[0])
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			this[index(0, j, cols)] = index(0, j, cols)
		}
		if board[len(board)-1][j] == 'O' {
			this[index(len(board)-1, j, cols)] = index(len(board)-1, j, cols)
		}
	}
	for i := 0; i < len(board); i++ {
		this[index(i, 0, cols)] = index(i, 0, cols)
		this[index(i, cols-1, cols)] = index(i, cols-1, cols)
	}
}