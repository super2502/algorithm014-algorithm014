package Day81

func countSmaller(nums []int) []int {
	ret := make([]int, len(nums))
	nodes := make([]*node, len(nums))
	for i := range nums {
		nodes[i] = &node{
			idx: i,
			val: nums[i],
		}
	}
	mergeSort(nodes, 0, len(nodes)-1)
	for _, node := range nodes {
		//fmt.Printf("%v:%v ", node.val, node.cnt)
		ret[node.idx] = node.cnt
	}
	return ret
}

func mergeSort(nodes []*node, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)>>1
	mergeSort(nodes, l, mid)
	mergeSort(nodes, mid+1, r)
	merge(nodes, l, mid, r)
}

func merge(nodes []*node, l, mid, r int) {
	tmp := make([]*node, r-l+1)
	i, j, k := mid, r, r-l
	for i >= l && j >= mid+1 {
		if nodes[i].val > nodes[j].val {
			nodes[i].cnt += j - mid
			tmp[k] = nodes[i]
			i--
		} else {
			tmp[k] = nodes[j]
			j--
		}
		k--
	}
	for i >= l {
		tmp[k] = nodes[i]
		k--
		i--
	}
	for j >= mid+1 {
		tmp[k] = nodes[j]
		k--
		j--
	}
	copy(nodes[l:r+1], tmp)
}

type node struct {
	idx int
	val int
	cnt int
}
