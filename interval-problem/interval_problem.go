// @file: interval_problem.go
// @date: 2021/1/18

// Package intervalproblem
package intervalproblem

import (
	"fmt"
	"sort"
)

type intervalList [][]int

func (i intervalList) Len() int {
	return len(i)
}

func (i intervalList) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i intervalList) Less(a, b int) bool {
	if i[a][0] == i[b][0] {
		return i[a][1] > i[b][1]
	}
	return i[a][0] < i[b][0]
}

func (i intervalList) String() string {
	res := ""
	for _, v := range i {
		res += fmt.Sprintf("(%d, %d) ", v[0], v[1])
	}
	return res
}

// 1288. 删除被覆盖区间
func removeCoveredIntervals(intervals [][]int) int {
	sort.Sort(intervalList(intervals))

	right := intervals[0][1]
	res := 0
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		// 排序之后interval[0]一定大于等于前一个区间的左侧，此时只需比较右侧
		if right >= interval[1] {
			res++ // 上个区间的右侧大于等于这个区间的右侧，则为覆盖区间
		} else {
			right = interval[1] // 否则更新右侧区间
		}
	}
	return len(intervals) - res
}

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(intervalList(intervals))

	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		current := intervals[i]
		if last[1] >= current[0] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			res = append(res, current)
		}
	}
	return res
}

// 986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	intervals := append(firstList, secondList...)
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(intervalList(intervals))
	res := make([][]int, 0)
	right := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if right > interval[1] {
			res = append(res, []int{interval[0], interval[1]})
		} else if right >= interval[0] && right <= interval[1] {
			res = append(res, []int{interval[0], right})
			right = interval[1]
		} else if right < interval[0] {
			right = interval[1]
		}
	}
	return res
}

func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	res := make([][]int, 0)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]
		// 存在交集
		if b2 >= a1 && a2 >= b1 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		if a2 > b2 {
			j++
		} else {
			i++
		}
	}
	return res
}
