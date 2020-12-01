package Week_16

import (
	"container/heap"
	"fmt"
)

func reorganizeString(S string) string {
	m := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		m[S[i]]++
	}
	nsHeap := make(nodes, 0)
	for b, cnt := range m {
		heap.Push(&nsHeap, &node{
			b:   b,
			cnt: cnt,
		})
	}
	//for _, node := range nsHeap {
	//	fmt.Printf("%+v\n", node)
	//}
	ret := ""
	for !nsHeap.IsEmpty() {
		first := heap.Pop(&nsHeap).(*node)
		fmt.Printf("first poped (%v)\n", string(first.b))
		if nsHeap.IsEmpty() {
			if first.cnt == 1 {
				ret += string(first.b)
				return ret
			} else {
				return ""
			}
		} else {
			second := heap.Pop(&nsHeap).(*node)
			fmt.Printf("second poped (%v)\n", string(second.b))
			ret += string(first.b) + string(second.b)
			if second.cnt > 1 {
				second.cnt--
				heap.Push(&nsHeap, second)
			}
		}
		if first.cnt > 1 {
			first.cnt--
			heap.Push(&nsHeap, first)
		}
	}
	return ret
}

type node struct {
	b   byte
	cnt int
}

type nodes []*node

func (ns *nodes) Len() int {
	return len(*ns)
}
func (ns *nodes) IsEmpty() bool {
	return ns.Len() == 0
}
func (ns *nodes) Swap(i, j int) {
	(*ns)[i], (*ns)[j] = (*ns)[j], (*ns)[i]
}
func (ns *nodes) Less(i, j int) bool {
	if (*ns)[i].cnt > (*ns)[j].cnt {
		return true
	} else if (*ns)[i].cnt < (*ns)[j].cnt {
		return false
	} else {
		return (*ns)[i].b < (*ns)[j].b
	}
}
func (ns *nodes) Push(x interface{}) {
	*ns = append(*ns, x.(*node))
}
func (ns *nodes) Pop() interface{} {
	x := (*ns)[ns.Len()-1]
	*ns = (*ns)[:ns.Len()-1]
	return x
}

var _ heap.Interface = &nodes{}
