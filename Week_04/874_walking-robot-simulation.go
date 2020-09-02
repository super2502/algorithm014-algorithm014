package Week_04

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	//-2：向左转 90 度
	//-1：向右转 90 度
	//1 <= x <= 9：向前移动 x 个单位长度
	obsMap := make(map[string]struct{})
	for _, obs := range obstacles {
		obsMap[fmt.Sprintf("%d.%d", obs[0], obs[1])] = struct{}{}
	}
	distX := []int{0, 1, 0, -1}
	distY := []int{1, 0, -1, 0}
	d := 0
	x, y := 0, 0

	maxArea := 0
	for _, command := range commands {
		if command == -1 {
			d = (d + 1) % 4
		} else if command == -2 {
			d = (d + 3) % 4
		} else {
			for k := 0; k < command; k++ {
				nextX := x + distX[d]
				nextY := y + distY[d]
				if _, ok := obsMap[fmt.Sprintf("%d.%d", nextX, nextY)]; ok {
					break
				}
				x = nextX
				y = nextY

				maxArea = max874(maxArea, x*x+y*y)
			}
		}
	}
	return maxArea
}

func max874(a, b int) int {
	if a > b {
		return a
	}
	return b
}
