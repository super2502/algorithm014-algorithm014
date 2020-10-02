package Week_08

import "sort"

type Leaderboard1 struct {
	playerScores [][]int
	playerMap    map[int][]int
}

func Constructor1() Leaderboard1 {
	return Leaderboard1{
		playerScores: make([][]int, 0),
		playerMap:    make(map[int][]int),
	}
}

func (this *Leaderboard1) AddScore(playerId int, score int) {
	ps, ok := this.playerMap[playerId]
	if !ok {
		ps = []int{playerId, score}
		this.playerScores = append(this.playerScores, ps)
		this.playerMap[playerId] = ps
	} else {
		ps[1] += score
	}
}

func (this *Leaderboard1) Top(K int) int {
	sort.Slice(this.playerScores, func(i, j int) bool {
		return this.playerScores[i][1] > this.playerScores[j][1]
	})
	total := 0
	if K > len(this.playerScores) {
		K = len(this.playerScores)
	}
	for k := 0; k < K; k++ {
		total += this.playerScores[k][1]
	}
	return total
}

func (this *Leaderboard1) Reset(playerId int) {
	ps, _ := this.playerMap[playerId]
	ps[1] = 0
}
