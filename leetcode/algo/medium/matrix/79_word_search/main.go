package wordsearch

// my solution (not fast)
var shift = [][]int{
	{-1, 0}, // down
	{1, 0},  // up
	{0, -1}, //left
	{0, 1},  // right
}

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	wordBytes := []byte(word)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var res bool
			if board[i][j] == wordBytes[0] {
				res = dfs(board, []int{i, j}, wordBytes, 0)
			}
			if res {
				// we found required word
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, point []int, word []byte, depth int) bool {
	// recursive Deep search with backtracking in case of invalid sequence
	if depth == len(word) {
		return true
	}
	i := point[0]
	j := point[1]
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[depth] {
		return false
	}

	tmp := board[i][j]
	board[i][j] = '.'                    // mark point as visited
	defer func() { board[i][j] = tmp }() // backtracking after checking all neighbors

	for _, s := range shift {
		if dfs(board, []int{i + s[0], j + s[1]}, word, depth+1) {
			// word found, no need to check next sequences
			return true
		}

	}
	return false // word not found
}

func padWithZeros(matrix [][]byte) [][]byte {
	rows := len(matrix)
	cols := len(matrix[0])

	padded := make([][]byte, rows+2)
	for i := range padded {
		padded[i] = make([]byte, cols+2)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			padded[i+1][j+1] = matrix[i][j]
		}
	}

	return padded
}

// not my solution
func existFaster(board [][]byte, word string) bool {
	visited := make([][]bool, 0, len(board))

	for range board {
		visited = append(visited, make([]bool, len(board[0])))
	}

	var dfs func(row, col, i int) bool

	dfs = func(i, j, depth int) bool {
		if visited[i][j] {
			return false
		}

		if board[i][j] != word[depth] {
			return false
		}

		if depth == len(word)-1 {
			// word found and previous check on == is ok
			return true
		}

		visited[i][j] = true

		// shift up
		if i > 0 && dfs(i-1, j, depth+1) {
			return true
		}

		// shift down
		if i < len(board)-1 && dfs(i+1, j, depth+1) {
			return true
		}

		// shift left
		if j > 0 && dfs(i, j-1, depth+1) {
			return true
		}

		// shift right
		if j < len(board[0])-1 && dfs(i, j+1, depth+1) {
			return true
		}

		visited[i][j] = false // backtracking

		return false
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] {
				found := dfs(i, j, 0)
				if found {
					return true
				}
			}

		}
	}

	return false
}
