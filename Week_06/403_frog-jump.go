package Week_06

func canCross(stones []int) bool {

	if len(stones) <= 1 {
		return true
	}
	stoneMap := make(map[int]map[int]struct{})
	for i := 1; i <len(stones); i++ {
		stoneMap[stones[i]] = make(map[int]struct{})
	}
	stoneMap[1] = map[int]struct{}{1: {}}
	for i := 1; i < len(stones); i++ {
		stoneIdx := stones[i]
		if len(stoneMap[stoneIdx]) == 0 {
			return false
		}
		for step := range stoneMap[stoneIdx] {
			for j := -1; j <= 1; j++ {
				jump := step + j
				if jump == 0 {
					continue
				}
				nextIdx := stoneIdx + jump
				if nextIdx == stones[len(stones) - 1] {
					return true
				}
				if _, ok := stoneMap[nextIdx]; !ok {
					continue
				}
				stoneMap[nextIdx][jump] = struct{}{}
			}
		}
	}
	if len(stoneMap[stones[len(stones)-1]]) > 0 {
		return true
	}
	return false
}