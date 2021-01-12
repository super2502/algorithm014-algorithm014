package Day115

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ret := make([]int, 0)
	for i := 0; i <= k; i++ {
		if i > len(nums1) || k-i > len(nums2) {
			continue
		}
		s1 := pick(nums1, len(nums1)-i)
		s2 := pick(nums2, len(nums2)-(k-i))
		//fmt.Printf("i: %v, s1 (%+v), s2(%+v)\n", i , s1, s2)
		m := merge(s1, s2)
		//fmt.Printf("compare (%+v)(%+v)", m, ret)
		if bigger(m, ret) {
			ret = m
		}
		//fmt.Printf(" and get (%+v)\n", ret)

	}
	return ret
}

func pick(nums []int, k int) []int {
	//fmt.Printf("pick in (%+v) (%v)\n", nums, k)
	s := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for len(s) > 0 && k > 0 && nums[i] > s[len(s)-1] {
			s = s[:len(s)-1]
			k--
		}
		s = append(s, nums[i])
	}
	for j := 0; j < k; j++ {
		s = s[:len(s)-1]
	}
	//fmt.Printf("pick out (%+v)\n", s)
	return s
}

func merge(a []int, b []int) []int {
	ret := make([]int, len(a)+len(b))
	for i := range ret {
		if bigger(a, b) {
			ret[i], a = a[0], a[1:]
		} else {
			ret[i], b = b[0], b[1:]
		}
	}
	return ret
}

func bigger(a []int, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return len(a) > len(b)
}
