package Week_01

func merge(nums1 []int, m int, nums2 []int, n int) {

	i, j := m-1, n-1
	for i >= -1 && j >= -1 {
		if i == -1 {
			for k := 0; k <= j; k++ {
				nums1[k] = nums2[k]
			}
			break
		}
		if j == -1 {
			break
		}

		if nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}

	}

}
