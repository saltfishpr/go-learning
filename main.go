package main

import "fmt"

type TimeZone struct {
	Name         string // 时区名称, 如: Asia/Shanghai
	DisplayName  string // 显示名称, 如: (UTC+08:00) Beijing
	Abbreviation string // 时区缩写, 如: CST
	StartTime    string // 时区开始时间, yyyy-MM-dd HH:mm:ss, 当地时区
	EndTime      string // 时区结束时间, yyyy-MM-dd HH:mm:ss, 当地时区
	Offset       int    // 偏移量, 单位 min
}

func (tz TimeZone) String() string {
	return fmt.Sprintf("%s: from: %s to: %s, displayName: %s", tz.Name, tz.StartTime, tz.EndTime, tz.DisplayName)
}

func main() {
	loader := &loader{
		downloadData: true,
		baseURL:      "https://timezonedb.com",
		fileName:     "TimeZoneDB.csv.zip",
	}

	zones, err := loader.Load()
	if err != nil {
		panic(err)
	}

	_ = zones
}
