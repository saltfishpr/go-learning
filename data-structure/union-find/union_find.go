// @file: union_find.go
// @date: 2021/2/3

// Package unionfind 并查集算法
package unionfind

type UF struct {
	count  int   // 连通分量个数
	parent []int // 节点x的根节点为parent[x]
	weight []int // 节点x为根的树的weight
}

func Constructor(n int) *UF {
	p := make([]int, n)
	w := make([]int, n)
	// 初始状态所有根节点都为自己
	for i := range p {
		p[i] = i
		w[i] = 1
	}
	return &UF{
		count:  n,
		parent: p,
		weight: w,
	}
}

// root 返回 x 的根节点
func (uf *UF) root(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]] // 进行路径压缩
		x = uf.parent[x]
	}
	return x
}

// Union 将两棵树合并为一棵
func (uf *UF) Union(p, q int) {
	rootP := uf.root(p)
	rootQ := uf.root(q)
	if rootP == rootQ {
		return
	}
	// 把小树接到大树下面
	if uf.weight[rootP] > uf.weight[rootQ] {
		uf.parent[rootQ] = rootP
		uf.weight[rootP] += uf.weight[rootQ]
	} else {
		uf.parent[rootP] = rootQ
		uf.weight[rootQ] += rootP
	}
	uf.count--
}

// connected 判断两个节点是否连通
func (uf *UF) connected(p, q int) bool {
	rootP := uf.root(p)
	rootQ := uf.root(q)
	return rootP == rootQ
}

// 130. 被围绕的区域
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	uf := Constructor(m*n + 1)
	dummy := m * n
	// 将四周的 'O' 与 dummy 相连
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(i*n, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.Union(i*n+n-1, dummy)
		}
	}
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			uf.Union(i, dummy)
		}
		if board[m-1][i] == 'O' {
			uf.Union((m-1)*n+i, dummy)
		}
	}

	// 将此O与上下左右的O连通
	d := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != 'O' {
				continue
			}
			for k := 0; k < 4; k++ {
				x, y := i+d[k][0], j+d[k][1]
				if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' {
					uf.Union(i*n+j, x*n+y)
				}
			}
		}
	}

	// 所有不与dummy连通的节点都被替换掉
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !uf.connected(i*n+j, dummy) {
				board[i][j] = 'X'
			}
		}
	}
}

// 990. 等式方程的可满足性
func equationsPossible(equations []string) bool {
	if len(equations) == 0 {
		return true
	}
	uf := Constructor(26)
	notEquals := make([]string, 0)
	var c byte = 'a'

	for _, s := range equations {
		a, b := s[0], s[3]
		if s[1:3] == "==" {
			uf.Union(int(a-c), int(b-c))
		} else if s[1:3] == "!=" {
			notEquals = append(notEquals, s)
		}
	}

	for _, s := range notEquals {
		a, b := s[0], s[3]
		if uf.connected(int(a-c), int(b-c)) {
			return false
		}
	}

	return true
}
