package main

import (
	"math"
)

func isRectangleCover(rectangles [][]int) bool {

	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64

	for _, rec := range rectangles {
		minX = min(minX, rec[0])
		minX = min(minX, rec[2])
		minY = min(minY, rec[1])
		minY = min(minY, rec[3])
		maxX = max(maxX, rec[0])
		maxX = max(maxX, rec[2])
		maxY = max(maxY, rec[1])
		maxY = max(maxY, rec[3])
	}

	//fmt.Printf("(%v)(%v)(%v)(%v)\n", minX, minY, maxX, maxY)

	area := make([][]int, maxX-minX)
	for i := 0; i < len(area); i++ {
		area[i] = make([]int, maxY-minY)
	}

	for _, rec := range rectangles {
		x0, x1 := rec[0], rec[2]
		if x0 > x1 {
			x0, x1 = x1, x0
		}
		y0, y1 := rec[1], rec[3]
		if y0 > y1 {
			y0, y1 = y1, y0
		}
		for x := x0; x < x1; x++ {
			for y := y0; y < y1; y++ {
				area[x-minX][y-minY]++
			}
		}

	}
	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {
			if area[i][j] != 1 {
				return false
			}
		}
	}
	return true

}
