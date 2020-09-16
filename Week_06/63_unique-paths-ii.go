package Week_06

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	} else {
		dp[0] = 0
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[j] = dp[j-1]
		} else {
			dp[j] = 0
		}
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[j] += dp[j-1]
			} else {
				dp[j] = 0
			}
		}
	}
	return dp[n-1]
}

func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	dp[0] = make([]int, n)
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		} else {
			dp[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = 0
		}

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}
