package DFS

/**
@author Shitanford
@create 2021-10-03 15:02
*/

func main() {

}

func exist(board [][]byte, word string) bool {
	selected := make([][]bool, len(board))
	for i := 0; i < len(selected); i++ {
		selected[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if helper(board, selected, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, selected [][]bool, x, y int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) ||
		selected[x][y] || board[x][y] != word[index] {
		return false
	}
	selected[x][y] = true
	existed := false
	existed = existed || helper(board, selected, x+1, y, word, index+1)
	existed = existed || helper(board, selected, x-1, y, word, index+1)
	existed = existed || helper(board, selected, x, y+1, word, index+1)
	existed = existed || helper(board, selected, x, y-1, word, index+1)
	selected[x][y] = false
	return existed
}