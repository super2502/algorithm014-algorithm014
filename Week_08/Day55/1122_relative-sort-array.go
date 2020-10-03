package Day55

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]int)
	for idx, num := range arr2 {
		arr2Map[num] = idx
	}

	sort.Slice(arr1, func(i, j int) bool {
		idxi, oki := arr2Map[arr1[i]]
		idxj, okj := arr2Map[arr1[j]]
		if oki && okj {
			return idxi < idxj // 都出现过，按照arr2 的index排
		}
		if oki {
			return true // i出现过，i放前边
		}
		if okj {
			return false // j出现过，j放前边
		}
		return arr1[i] < arr1[j] // 都没出现过，按大小排
	})

	return arr1
}
