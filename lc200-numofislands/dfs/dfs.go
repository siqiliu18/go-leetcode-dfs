package dfs

func DFS(x int, y int, rowLen int, colLen int, grid [][]byte, visited [][]bool) {

	// left
	if y-1 >= 0 && string(grid[x][y-1]) == "1" && !visited[x][y-1] {
		visited[x][y-1] = true
		DFS(x, y-1, rowLen, colLen, grid, visited)
	}

	// right
	if y+1 < colLen && string(grid[x][y+1]) == "1" && !visited[x][y+1] {
		visited[x][y+1] = true
		DFS(x, y+1, rowLen, colLen, grid, visited)
	}

	// top
	if x-1 >= 0 && string(grid[x-1][y]) == "1" && !visited[x-1][y] {
		visited[x-1][y] = true
		DFS(x-1, y, rowLen, colLen, grid, visited)
	}

	// bottom
	if x+1 < rowLen && string(grid[x+1][y]) == "1" && !visited[x+1][y] {
		visited[x+1][y] = true
		DFS(x+1, y, rowLen, colLen, grid, visited)
	}
}

func NumIslands(grid [][]byte) int {
	res := 0

	rowLen := len(grid)
	colLen := len(grid[0])

	visited := make([][]bool, rowLen)

	for i := 0; i < rowLen; i++ {
		visited[i] = make([]bool, colLen)
		for j := 0; j < colLen; j++ {
			visited[i][j] = false
		}
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if string(grid[i][j]) == "1" && !visited[i][j] {
				DFS(i, j, rowLen, colLen, grid, visited)
				res++
			}
		}
	}

	return res
}