package minimum_path_sum

func minPathSum(grid [][]int) int {
	cols := len(grid[0])
	rows := len(grid)
	result := make([][]int, rows)
	prevSum := 0

	copy(result, grid)

	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)

		result[i][0] = prevSum + grid[i][0]
		prevSum = result[i][0]
	}

	prevSum = 0

	for i := 0; i < cols; i++ {
		result[0][i] = prevSum + grid[0][i]
		prevSum = result[0][i]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			result[i][j] = min(result[i-1][j]+grid[i][j], result[i][j-1]+grid[i][j])
		}
	}

	return result[rows-1][cols-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
