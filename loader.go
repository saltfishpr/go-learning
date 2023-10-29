package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cast"
)

type loader struct {
	downloadData bool
	baseURL      string // https://timezonedb.com
	fileName     string // TimeZoneDB.csv.zip
}

func (l *loader) Load() ([]*TimeZone, error) {
	const dataDir = "data"

	if l.downloadData {
		if err := l.prepareData(dataDir); err != nil {
			return nil, err
		}
	}

	zones, err := l.load(dataDir)
	if err != nil {
		return nil, err
	}

	keys := lo.Keys(zones)
	sort.Strings(keys)

	var res []*TimeZone
	for _, key := range keys {
		zone := zones[key]
		for _, trans := range zone.Trans {
			cityName, err := lo.Last(strings.Split(zone.Name, "/"))
			if err != nil {
				return nil, err
			}
			displayName := fmt.Sprintf("(UTC%s) %s", getTimeZoneOffsetString(trans.Offset), cityName)

			startTime := time.Unix(trans.StartTime, 0).UTC().Add(time.Duration(trans.Offset) * time.Second).Format(time.DateTime)
			var endTime string
			if trans.EndTime != nil {
				endTime = time.Unix(*trans.EndTime, 0).UTC().Add(time.Duration(trans.Offset) * time.Second).Format(time.DateTime)
			}
			res = append(res, &TimeZone{
				Name:         zone.Name,
				DisplayName:  displayName,
				Abbreviation: trans.Abbreviation,
				StartTime:    startTime,
				EndTime:      endTime,
				Offset:       trans.Offset / 60,
			})
		}
	}
	return res, nil
}

func getTimeZoneOffsetString(offset int) string {
	if offset < 0 {
		offset = -offset
		return fmt.Sprintf("-%02d:%02d", offset/3600, (offset%3600)/60)
	}
	return fmt.Sprintf("+%02d:%02d", offset/3600, (offset%3600)/60)
}

func (l *loader) prepareData(dir string) error {
	downloadPath := path.Join(os.TempDir(), l.fileName)
	if _, err := os.Stat(downloadPath); err == nil {
		log.Println("old time zone db file found, removing")
		if err := os.Remove(downloadPath); err != nil {
			return fmt.Errorf("remove old time zone db file error: %w", err)
		}
	}

	resp, err := http.Get(l.baseURL + "/files/" + l.fileName)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(downloadPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("create temp file error: %w", err)
	}
	defer f.Close()

	log.Printf("downloading time zone db file to %s\n", downloadPath)
	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading")
	if _, err := io.Copy(io.MultiWriter(f, bar), resp.Body); err != nil {
		return fmt.Errorf("download error: %w", err)
	}

	cmd := exec.Command("unzip", "-o", "-d", dir, downloadPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("unzip error: %w", err)
	}

	return nil
}

type loaderZone struct {
	Name  string
	Trans []*loaderZoneTran
}

type loaderZoneTran struct {
	Abbreviation string // 时区缩写
	StartTime    int64  // 开始时间, 单位 s
	EndTime      *int64 // 开始时间, 单位 s, nil 表示没有结束时间
	Offset       int    // 偏移量, 单位 s
	IsDST        bool   // 是否是夏令时
}

func (l *loader) load(dataDir string) (map[string]*loaderZone, error) {
	const (
		countryFile  = "country.csv"
		timeZoneFile = "time_zone.csv"
	)

	type countryColumn int
	const (
		countryCode countryColumn = iota
		countryName
	)

	type timeZoneColumn int
	const (
		timeZoneName timeZoneColumn = iota
		timeZoneCountryCode
		timeZoneAbbreviation
		timeZoneStartTime
		timeZoneOffset
		timeZoneIsDST
	)

	zones := make(map[string]*loaderZone, 10)
	f, err := os.Open(path.Join(dataDir, timeZoneFile))
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)

	var lastZoneTran *loaderZoneTran
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		zone, ok := zones[row[timeZoneName]]
		if !ok {
			zone = &loaderZone{
				Name: row[timeZoneName],
			}
			zones[row[timeZoneName]] = zone
		}
		zoneTran := &loaderZoneTran{
			Abbreviation: row[timeZoneAbbreviation],
			StartTime:    cast.ToInt64(row[timeZoneStartTime]),
			EndTime:      nil,
			Offset:       cast.ToInt(row[timeZoneOffset]),
			IsDST:        row[timeZoneIsDST] == "1",
		}
		zone.Trans = append(zone.Trans, zoneTran)

		if lastZoneTran != nil {
			lastEndTime := zoneTran.StartTime - 1
			lastZoneTran.EndTime = &lastEndTime
		}
		lastZoneTran = zoneTran
	}

	return zones, nil
}
