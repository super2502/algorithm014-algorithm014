package Week_08

func reversePairs(nums []int) int {
	c := mergeSortAndCount(nums, 0, len(nums)-1)
	//fmt.Printf("%+v\n", nums)
	return c
}

func mergeSortAndCount(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)>>1
	c1 := mergeSortAndCount(nums, left, mid)
	c2 := mergeSortAndCount(nums, mid+1, right)
	c3 := count493(nums, left, mid, right)
	merge493(nums, left, mid, right)
	return c1 + c2 + c3
}

func merge493(nums []int, left int, mid int, right int) {
	tmp := make([]int, right-left+1)
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			tmp[k] = nums[i]
			k++
			i++
		} else {
			tmp[k] = nums[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmp[k] = nums[i]
		k++
		i++
	}
	for j <= right {
		tmp[k] = nums[j]
		k++
		j++
	}
	copy(nums[left:right+1], tmp)
}

func count493(nums []int, left, mid, right int) int {
	cnt := 0
	i, j := left, mid+1
	for i <= mid && j <= right {
		if nums[i] <= 2*nums[j] {
			i++
		} else {
			cnt += mid - i + 1
			j++
		}
	}
	return cnt
}
