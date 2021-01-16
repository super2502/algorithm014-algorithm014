package Day158

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	visited[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		tmp := make([]int, 0)
		for _, idx := range queue {
			if arr[idx] == 0 {
				return true
			}
			left := idx - arr[idx]
			right := idx + arr[idx]
			if left >= 0 && !visited[left] {
				visited[left] = true
				tmp = append(tmp, left)
			}
			if right < len(arr) && !visited[right] {
				visited[right] = true
				tmp = append(tmp, right)
			}
		}
		queue = tmp
	}
	return false
}
