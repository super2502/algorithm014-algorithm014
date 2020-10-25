package Week_10

import (
	"container/heap"
)

func rearrangeString(s string, k int) string {
	if k == 0 {
		return s
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	pq := &priorityQueue{}
	for b, cnt := range m {
		node := &node358{b: b, cnt: cnt}
		heap.Push(pq, node)
	}
	ret := ""
	for pq.Len() > 0 {
		tmp := make([]*node358, 0)
		tmp1 := make([]*node358, 0)
		for i := 0; i < k; i++ {
			if pq.Len() == 0 {
				if len(tmp1) == 0 {
					break
				} else {
					return ""
				}
			}
			node := heap.Pop(pq).(*node358)
			//ret += string(node.b)
			node.cnt--
			tmp = append(tmp, node)
			if node.cnt > 0 {
				tmp1 = append(tmp1, node)
			}
		}
		for _, node := range tmp {
			ret += string(node.b)
		}
		for _, node := range tmp1 {
			heap.Push(pq, node)
		}
	}

	return ret
}

type node358 struct {
	b   byte
	cnt int
}

type priorityQueue []*node358

func (pq *priorityQueue) Len() int {
	return len(*pq)
}
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *priorityQueue) Less(i, j int) bool { // 大顶堆
	if (*pq)[i].cnt > (*pq)[j].cnt {
		return true
	} else if (*pq)[i].cnt < (*pq)[j].cnt {
		return false
	} else {
		return (*pq)[i].b < (*pq)[j].b
	}
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node358))
}
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}

var _ heap.Interface = &priorityQueue{}
