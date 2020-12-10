package Day123

func moveZeroes(nums []int) {
	idx := 0
	for i, num := range nums {
		if num != 0 {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	//fmt.Printf("%+v\n", nums)
}
