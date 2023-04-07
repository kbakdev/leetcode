package _go

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)
		dfs(grid, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(grid, 0, j)
		dfs(grid, m-1, j)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i int, i2 int) {
	if i < 0 || i >= len(grid) || i2 < 0 || i2 >= len(grid[0]) || grid[i][i2] == 0 {
		return
	}
	grid[i][i2] = 0
	dfs(grid, i+1, i2)
	dfs(grid, i-1, i2)
	dfs(grid, i, i2+1)
	dfs(grid, i, i2-1)
}
