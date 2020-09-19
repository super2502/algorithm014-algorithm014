package Day41

// [4,5,6,7,0,1,2]
// [5,6,7,0,1,2,3,4]
// [4,5,6,7,8,0,1,2]
func search(nums []int, target int) int {

	l, r := 0, len(nums) - 1
	for l <= r {
		mid := l + ( r - l ) / 2
		if nums[mid] == target {
			return mid
		} else if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[l] > nums[mid] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
