package main

func searchLastLessEq(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + ( r - l ) >> 1
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func searchFirstLargeEq(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + ( r - l ) >> 1
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func searchFirstEq(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + ( r - l ) >> 1
		if nums[mid] == target {
			if mid > 0 && nums[mid] == nums[mid-1] {
				r = mid - 1
			} else {
				return mid
			}
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func searchLastEq(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + ( r - l ) >> 1
		if nums[mid] == target {
			if mid < len(nums) - 1 && nums[mid] == nums[mid+1] {
				l = mid + 1
			} else {
				return mid
			}
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}