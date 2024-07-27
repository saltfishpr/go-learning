package main

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		visited[v] = true
	}
	if visited["0000"] {
		return -1
	}

	queue := [][]byte{{'0', '0', '0', '0'}}
	depth := 0
	for len(queue) != 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			current := queue[i]
			if string(current) == target {
				return depth
			}
			for j := 0; j < len(current); j++ {
				turnUp(current, j)
				if !visited[string(current)] {
					visited[string(current)] = true
					currentCopy := make([]byte, len(current))
					copy(currentCopy, current)
					queue = append(queue, currentCopy)
				}
				turnDown(current, j)

				turnDown(current, j)
				if !visited[string(current)] {
					visited[string(current)] = true
					currentCopy := make([]byte, len(current))
					copy(currentCopy, current)
					queue = append(queue, currentCopy)
				}
				turnUp(current, j)
			}
		}
		depth++
		queue = queue[queueSize:]
	}

	return -1
}

func turnUp(lock []byte, idx int) {
	lock[idx] = lock[idx] + 1
	if lock[idx] > '9' {
		lock[idx] = '0'
	}
}

func turnDown(lock []byte, idx int) {
	lock[idx] = lock[idx] - 1
	if lock[idx] < '0' {
		lock[idx] = '9'
	}
}
