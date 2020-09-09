package Day31

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	obsMap := make(map[string]struct{})
	for i := 0; i < len(obstacles); i++ {
		obsMap[fmt.Sprintf("%d.%d", obstacles[i][0], obstacles[i][1])] = struct{}{}
	}
	x, y := 0, 0
	dictX := []int{0, 1, 0, -1}
	dictY := []int{1, 0, -1, 0}
	p := 0
	maxDist := 0
	for _, cmd := range commands {
		switch cmd {
		case -1:
			p = (p + 1) % 4
		case -2:
			p = (p + 3) % 4
		default:
			for i := 0; i < cmd; i++ {
				nextX := x + dictX[p]
				nextY := y + dictY[p]
				if _, ok := obsMap[fmt.Sprintf("%d.%d", nextX, nextY)]; ok {
					break
				}
				x = nextX
				y = nextY
				maxDist = max(maxDist, x*x+y*y)
			}

		}
	}
	return maxDist
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
