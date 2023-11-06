package internal

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cast"
)

var defaultLoaderOptions = &LoaderOptions{
	baseURL:  "https://timezonedb.com",
	fileName: "TimeZoneDB.csv.zip",
}

type LoaderOptions struct {
	baseURL  string
	fileName string
}

type LoaderOption func(*LoaderOptions)

func WithBaseURL(baseURL string) LoaderOption {
	return func(o *LoaderOptions) {
		o.baseURL = baseURL
	}
}

func WithFileName(fileName string) LoaderOption {
	return func(o *LoaderOptions) {
		o.fileName = fileName
	}
}

func evaluateLoaderOptions(opts ...LoaderOption) LoaderOptions {
	// default options
	options := *defaultLoaderOptions
	for _, o := range opts {
		o(&options)
	}
	return options
}

type loader struct {
	dataDir  string
	baseURL  string
	fileName string
}

func NewLoader(dataDir string, opts ...LoaderOption) *loader {
	opt := evaluateLoaderOptions(opts...)
	return &loader{
		dataDir:  dataDir,
		baseURL:  opt.baseURL,
		fileName: opt.fileName,
	}
}

func (l *loader) Load(ctx context.Context, downloadData bool) ([]*TimeZone, error) {
	if downloadData {
		if err := l.prepareData(ctx, l.dataDir); err != nil {
			return nil, err
		}
	}

	zones, err := l.loadFromDir(l.dataDir)
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

			startTime := TimestampToTimeString(trans.StartTime, trans.Offset/60)
			var endTime string
			if trans.EndTime != nil {
				endTime = TimestampToTimeString(*trans.EndTime, trans.Offset/60)
			} else {
				endTime = MaxTimeString
			}
			res = append(res, &TimeZone{
				Name:         zone.Name,
				DisplayName:  displayName,
				Abbreviation: trans.Abbreviation,
				StartTime:    startTime,
				EndTime:      endTime,
				IsDST:        trans.IsDST,
				Offset:       trans.Offset / 60,
				Source:       lo.Must(url.Parse(l.baseURL)).Host,
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

func (l *loader) prepareData(ctx context.Context, dir string) error {
	downloadPath := path.Join(os.TempDir(), l.fileName)
	if _, err := os.Stat(downloadPath); err == nil {
		log.Println("old timezone db file found, removing")
		if err := os.Remove(downloadPath); err != nil {
			return fmt.Errorf("remove old timezone db file error: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.baseURL+"/files/"+l.fileName, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(downloadPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("create temp file error: %w", err)
	}
	defer f.Close()

	log.Printf("downloading timezone db file to %s\n", downloadPath)
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

func (l *loader) loadFromDir(dataDir string) (map[string]*loaderZone, error) {
	const timeZoneFile = "time_zone.csv"
	f, err := os.Open(path.Join(dataDir, timeZoneFile))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parseTimeZones(f)
}

// parseTimeZones parse timezone csv file.
func parseTimeZones(reader io.Reader) (map[string]*loaderZone, error) {
	type timeZoneColumn int
	const (
		timeZoneName timeZoneColumn = iota
		timeZoneCountryCode
		timeZoneAbbreviation
		timeZoneStartTime
		timeZoneOffset
		timeZoneIsDST
	)

	r := csv.NewReader(reader)

	zones := make(map[string]*loaderZone, 10)
	var lastZoneName string
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

		if lastZoneTran != nil && lastZoneName == zone.Name {
			lastEndTime := zoneTran.StartTime
			lastZoneTran.EndTime = &lastEndTime
		}
		lastZoneName = zone.Name
		lastZoneTran = zoneTran
	}

	return zones, nil
}
