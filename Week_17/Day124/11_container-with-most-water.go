package Day124

func maxArea(height []int) int {
	ret := 0
	i, j := 0, len(height)-1
	for i < j {
		ret = max(ret, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] {
			for i < j && height[i] > height[i+1] {
				i++
			}
			i++
		} else {
			for i < j && height[j] > height[j-1] {
				j--
			}
			j--
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
