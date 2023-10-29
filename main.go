package main

type Zone struct {
	Name  string
	Trans []*ZoneTran
}

type ZoneTran struct {
	Abbreviation string // 时区缩写
	StartUTC     int64  // 开始时间, 单位 s
	Offset       int    // 偏移量, 单位 s
	IsDST        bool   // 是否是夏令时
}

func main() {
	loader := &loader{
		baseURL:  "https://timezonedb.com",
		fileName: "TimeZoneDB.csv.zip",
	}

	zones, err := loader.Load()
	if err != nil {
		panic(err)
	}

	_ = zones
}
