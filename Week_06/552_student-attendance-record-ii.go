package Week_06

func checkRecord(n int) int {
	if n == 0 {
		return 0
	}
	dp00 := 1
	dp01 := 1
	dp02 := 0
	dp10 := 1
	dp11 := 0
	dp12 := 0

	for i := 1; i <= n; i++ {
		dp12Tmp, dp11Tmp := dp12, dp11
		dp02Tmp, dp01Tmp := dp02, dp01
		dp12 = dp11
		dp11 = dp10
		dp10 = (dp10 + dp11Tmp + dp12Tmp + dp00 + dp01 + dp02) % (1e9 + 7)
		dp02 = dp01
		dp01 = dp00
		dp00 = (dp00 + dp01Tmp + dp02Tmp) % (1e9 + 7)
	}
	return dp10
}

func checkRecord1(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 2)
		dp[i][0] = make([]int, 3)
		dp[i][1] = make([]int, 3)
	}
	dp[0][0][0] = 1
	dp[0][0][1] = 1
	dp[0][0][2] = 0
	dp[0][1][0] = 1
	dp[0][1][1] = 0
	dp[0][1][2] = 0

	for i := 1; i < n; i++ {
		dp[i][0][0] = (dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % (1e9 + 7)
		dp[i][0][1] = dp[i-1][0][0]
		dp[i][0][2] = dp[i-1][0][1]
		dp[i][1][0] = (dp[i-1][1][0] + dp[i-1][1][1] + dp[i-1][1][2] + dp[i-1][0][0] + dp[i-1][0][1] + dp[i-1][0][2]) % (1e9 + 7)
		dp[i][1][1] = dp[i-1][1][0]
		dp[i][1][2] = dp[i-1][1][1]
	}

	x := dp[n-1]
	return (x[0][0] + x[0][1] + x[0][2] + x[1][0] + x[1][1] + x[1][2]) % (1e9 + 7)
}
