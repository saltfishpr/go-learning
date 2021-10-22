// @file: queue.go
// @date: 2021/1/23

// Package stackandqueue
package stackandqueue

type queueInt struct {
	length int
	queue  []int
}

func (q *queueInt) push(x int) {
	q.queue = append(q.queue, x)
	q.length++
}

func (q *queueInt) pop() int {
	if q.length == 0 {
		return 0
	}
	res := q.queue[0]
	q.queue = q.queue[1:]
	q.length--
	return res
}

func (q *queueInt) peek() int {
	if q.length == 0 {
		return 0
	}
	return q.queue[0]
}

// 542. 01 矩阵
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	rows := len(matrix)
	cols := len(matrix[0])

	queue := make([][]int, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				queue = append(queue, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}

	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		for _, v := range directions {
			x := point[0] + v[0]
			y := point[1] + v[1]
			if x >= 0 && x < rows && y >= 0 && y < cols && matrix[x][y] == -1 {
				queue = append(queue, []int{x, y})
				matrix[x][y] = matrix[point[0]][point[1]] + 1
			}
		}
	}
	return matrix
}
