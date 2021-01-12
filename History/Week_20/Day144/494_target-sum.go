package Day144

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < S {
		return 0
	}
	S += sum
	if S%2 == 1 {
		return 0
	}
	S /= 2
	dp := make([][]int, len(nums)+1)
	dp[0] = make([]int, S+1)
	dp[0][0] = 1
	for i := 1; i <= len(nums); i++ {
		dp[i] = make([]int, S+1)
		for j := 0; j <= S; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
		}
	}
	//for i := 0; i<= len(nums); i++ {
	//	fmt.Printf("%+v\n", dp[i])
	//}
	return dp[len(nums)][S]
}

func findTargetSumWays1(nums []int, S int) int {
	ret := 0
	var dfs func(i int, sum int)
	dfs = func(i int, sum int) {
		if i == len(nums) {
			if sum == 0 {
				ret++
			}
			return
		}
		dfs(i+1, sum-nums[i])
		dfs(i+1, sum+nums[i])
	}
	dfs(0, S)
	return ret
}
