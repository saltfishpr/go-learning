package internal

// TimeZone 时区信息, 在 [StartTime, EndTime) 时间段内, 该地区的时区偏移量为 Offset 分钟.
type TimeZone struct {
	ID           string `json:"id"`           // 时区 ID, readOnly
	Name         string `json:"name"`         // 时区名称, 如: Asia/Shanghai
	DisplayName  string `json:"displayName"`  // 显示名称, 如: (UTC+08:00) Beijing
	Abbreviation string `json:"abbreviation"` // 时区缩写, 如: CST
	StartTime    string `json:"startTime"`    // 时区开始时间, yyyy-MM-dd HH:mm:ss, 当地时区
	EndTime      string `json:"endTime"`      // 时区结束时间, yyyy-MM-dd HH:mm:ss, 当地时区
	IsDST        bool   `json:"isDST"`        // 是否为夏令时
	Offset       int    `json:"offset"`       // 偏移量, 单位 min
	Source       string `json:"source"`       // 数据来源, 如: timezonedb.com
}

func (tz TimeZone) Contains(s string) bool {
	return tz.StartTime <= s && (tz.EndTime == MaxTimeString || tz.EndTime > s)
}

func (tz TimeZone) Convert(s string) (int64, error) {
	ts, err := TimeString2Timestamp(s, tz.Offset)
	if err != nil {
		return 0, err
	}
	return ts, nil
}
