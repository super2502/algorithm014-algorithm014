package Week_08

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]int)
	for idx, num := range arr2 {
		arr2Map[num] = idx
	}
	sort.Slice(arr1, func(i, j int) bool {
		idxI, okI := arr2Map[arr1[i]]
		idxJ, okJ := arr2Map[arr1[j]]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
