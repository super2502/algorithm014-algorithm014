package Week_10

//var cnt int
func numberOfPatterns(m int, n int) int {
	ret := 0
	//for k :=0; k< 9;k++ {
	visited := [9]bool{}
	ret += count(m, n, 1, 0, visited) * 4
	visited = [9]bool{}
	ret += count(m, n, 1, 1, visited) * 4
	visited = [9]bool{}
	ret += count(m, n, 1, 4, visited)
	//}
	//fmt.Printf("i: %v, cnt: %v\n", i, cnt)
	return ret
}

func count(m, n int, level int, i int, visited [9]bool) int {
	ret := 0
	if level > n {
		return 0
	}
	if level >= m {
		ret++
	}
	visited[i] = true
	xi, yi := i/3, i%3
	for j := 0; j < 9; j++ {
		if j == i {
			continue
		}
		if visited[j] {
			continue
		}
		xj, yj := j/3, j%3
		if xi == xj && abs(yi-yj) == 2 {
			mid := xi*3 + (yi+yj)/2
			if !visited[mid] {
				continue
			}
		}
		if yi == yj && abs(xi-xj) == 2 {
			mid := (xi+xj)/2*3 + yi
			if !visited[mid] {
				continue
			}
		}
		if abs(yi-yj) == 2 && abs(xi-xj) == 2 {
			if !visited[4] {
				continue
			}
		}
		ret += count(m, n, level+1, j, visited)
	}
	visited[i] = false
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
