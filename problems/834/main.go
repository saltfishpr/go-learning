package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	sz := make([]int, n) // 存储子节点数（包括自己）
	dp := make([]int, n) // dp[i] 表示以 i 为根的树，所有子节点到它的距离之和
	var dfs func(int, int)
	// u 当前节点，f 父节点
	dfs = func(u, f int) {
		sz[u] = 1
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			dfs(v, u)
			dp[u] += dp[v] + sz[v]
			sz[u] += sz[v]
		}
	}
	dfs(0, -1)

	res := make([]int, n) // 存储最终结果
	var dfs2 func(int, int)
	dfs2 = func(u, f int) {
		res[u] = dp[u]
		for _, v := range graph[u] {
			if v == f {
				continue
			}
			// 存下当前结果
			pu, pv := dp[u], dp[v]
			su, sv := sz[u], sz[v]

			// 换根
			dp[u] -= dp[v] + sz[v]
			sz[u] -= sz[v]
			dp[v] += dp[u] + sz[u]
			sz[v] += sz[u]

			// 当前状态，v 已经为根节点，再进一步将 v 的子节点作为根节点，计算 ans
			dfs2(v, u)

			// 恢复
			dp[u], dp[v] = pu, pv
			sz[u], sz[v] = su, sv
		}
	}

	dfs2(0, -1)
	return res
}
