package Day121

func majorityElement(nums []int) int {
	A := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			A = nums[i]
			cnt = 1
			continue
		}
		if nums[i] != A {
			cnt--
			continue
		}
		cnt++
	}
	return A
}
