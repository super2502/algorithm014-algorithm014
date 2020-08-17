package Day8

func intersect(nums1 []int, nums2 []int) []int {

	ret := make([]int, 0)
	m1 := make(map[int]int)

	for _, num1 := range nums1 {
		m1[num1]++
	}
	for _, num2 := range nums2 {
		if m1[num2] > 0 {
			ret = append(ret, num2)
			m1[num2]--
		}
	}
	return ret
}
