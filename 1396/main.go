package main

import "strings"

type CheckInInfo struct {
	Station string
	Time    int
}

type Cost struct {
	Sum int // 总耗时
	Num int // 人数
}

type UndergroundSystem struct {
	check map[int]CheckInInfo
	data  map[string]Cost
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		check: make(map[int]CheckInInfo),
		data:  make(map[string]Cost),
	}
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.check[id] = CheckInInfo{Station: stationName, Time: t}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkInInfo := us.check[id]
	travel := strings.Join([]string{checkInInfo.Station, stationName}, "->")
	oldCost := us.data[travel]
	us.data[travel] = Cost{
		Sum: oldCost.Sum + t - checkInInfo.Time,
		Num: oldCost.Num + 1,
	}
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	travel := strings.Join([]string{startStation, endStation}, "->")
	if cost, ok := us.data[travel]; ok {
		return float64(cost.Sum) / float64(cost.Num)
	}
	return 0
}
