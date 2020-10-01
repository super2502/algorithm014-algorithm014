package Week_08

import "fmt"

func bubbleSort(nums []int) {
	loop := 0
	swap := 0
	for i := 0; i < len(nums); i++ {
		done := true
		for j := 1; j < len(nums)-i; j++ {
			loop++
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				done = false
				swap++
			}
		}
		if done {
			break
		}
	}
	fmt.Printf("loop (%v) times, swap (%v) times\n", loop, swap)
}
