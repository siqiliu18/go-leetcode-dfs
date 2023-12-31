package dfs

import (
	"math"
)

type Coordi struct {
	x   int
	y   int
	val int
}

func ShortestDistance(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	grid2 := make([][]int, rows)

	ones := []Coordi{}

	for x := 0; x < rows; x++ {
		grid2[x] = make([]int, cols)
		for y := 0; y < cols; y++ {
			grid2[x][y] = 0
			if grid[x][y] == 1 {
				ones = append(ones, Coordi{x, y, 1})
			}
		}
	}

	res := math.MaxInt
	round := 0
	for i, oneCoordi := range ones {
		round = i * -1
		// res = bfs(oneCoordi, grid, grid2, rows, cols, round)
		bfs(oneCoordi, grid, grid2, rows, cols, round)
	}

	// if res = bfs(oneCoordi, grid, grid2, rows, cols, round) above, no need to do this step
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] != 1 && grid[x][y] != 2 && grid[x][y] == round-1 {
				if grid2[x][y] < res {
					res = grid2[x][y]
				}
			}
		}
	}

	if res == math.MaxInt {
		return -1
	}

	return res
}

// currSum: as starting over more 1s, the values in grid2 will increase
// currSum stores the sum of steps to increasing 1s
func bfs(oneCoordi Coordi, grid, grid2 [][]int, rows, cols, round int) int {
	fourD := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Coordi{oneCoordi}

	currSum := math.MaxInt
	for len(queue) != 0 {
		currCoor := queue[0]
		queue = queue[1:]

		for _, di := range fourD {
			if currCoor.x+di[0] >= 0 && currCoor.x+di[0] < rows && currCoor.y+di[1] >= 0 && currCoor.y+di[1] < cols {
				// cannot use grid[currCoor.x+di[0]][currCoor.y+di[1]] != round here because != will visit all empty lands even two buildings are blocked and unreachable from each other
				if grid[currCoor.x+di[0]][currCoor.y+di[1]] == round {
					grid2[currCoor.x+di[0]][currCoor.y+di[1]] = currCoor.val + grid2[currCoor.x+di[0]][currCoor.y+di[1]]
					grid[currCoor.x+di[0]][currCoor.y+di[1]] -= 1
					queue = append(queue, Coordi{currCoor.x + di[0], currCoor.y + di[1], currCoor.val + 1})
					currSum = min(currSum, grid2[currCoor.x+di[0]][currCoor.y+di[1]])
				}
			}
		}
	}
	return currSum
}

// hint and code from https://www.youtube.com/watch?v=yjHXS2w_IvY
func ShortestDistance1(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	dist_matrix := make([][]int, rows)
	for x := 0; x < rows; x++ {
		dist_matrix[x] = make([]int, cols)
		for y := 0; y < cols; y++ {
			dist_matrix[x][y] = 0
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	BUILDING := 1
	// OBSTACLE := 2
	EMPTY_LAND := 0

	min_dist := math.MaxInt

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == BUILDING {
				local_dist := math.MaxInt
				queue := []Coordi{{row, col, 0}}
				for len(queue) != 0 {
					curr := queue[0]
					queue = queue[1:]
					cur_row := curr.x
					cur_col := curr.y
					distance := curr.val
					for _, di := range directions {
						new_row := cur_row + di[0]
						new_col := cur_col + di[1]
						if new_row >= 0 && new_row < rows && new_col >= 0 && new_col < cols && grid[new_row][new_col] == EMPTY_LAND {
							grid[new_row][new_col] -= 1
							dist_matrix[new_row][new_col] += (distance + 1)
							queue = append(queue, Coordi{new_row, new_col, distance + 1})
							local_dist = min(local_dist, dist_matrix[new_row][new_col])
						}
					}
				}
				EMPTY_LAND -= 1
				min_dist = local_dist
			}
		}
	}

	if min_dist == math.MaxInt {
		return -1
	}

	return min_dist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
