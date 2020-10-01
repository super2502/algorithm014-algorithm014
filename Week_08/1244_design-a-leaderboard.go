package Week_08

import (
	"container/heap"
	"math"
)

type Leaderboard struct {
	pq    priorityQueue
	cache map[int]*playerScore
}

func Constructor1244() Leaderboard {
	pq := make(priorityQueue, 0)
	return Leaderboard{
		pq:    pq,
		cache: make(map[int]*playerScore),
	}
}

func (this *Leaderboard) AddScore(playerId int, score int) {
	ps, ok := this.cache[playerId]
	if ok {
		ps.score += score
		heap.Init(&this.pq)
	} else {
		ps = &playerScore{
			id:    playerId,
			score: score,
		}
		heap.Push(&this.pq, ps)
		this.cache[playerId] = ps
	}
}

func (this *Leaderboard) Top(K int) int {
	tmp := make(priorityQueue, this.pq.Len())
	copy(tmp, this.pq)
	i := 0
	scores := 0
	for tmp.Len() > 0 && i < K {
		scores += heap.Pop(&tmp).(*playerScore).score
		i++
	}
	return scores
}

func (this *Leaderboard) Reset(playerId int) {
	ps, _ := this.cache[playerId]
	ps.score = math.MaxInt64
	heap.Init(&this.pq)
	heap.Pop(&this.pq)
	delete(this.cache, playerId)
}

type playerScore struct {
	id    int
	score int
}
type priorityQueue []*playerScore

func (pq *priorityQueue) Len() int {
	return len(*pq)
}
func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].score > (*pq)[j].score // 大顶堆
}
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return x
}
func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*playerScore))
}

var _ heap.Interface = &priorityQueue{}
