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

	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cast"
)

type Country struct {
	Code  string
	Name  string
	Zones []*Zone
}

type Zone struct {
	Name      string
	Countries []*Country
	Trans     []*ZoneTran
}

type ZoneTran struct {
	Abbreviation string // 时区缩写
	StartUTC     int64  // 开始时间
	Offset       int    // 偏移量
	IsDST        bool   // 是否是夏令时
}

func main() {
	if err := prepareData(); err != nil {
		log.Fatal(err)
	}

	countries, err := loadCountries()
	if err != nil {
		log.Fatal(err)
	}

	tz, err := loadTimeZones(countries)
	if err != nil {
		log.Fatal(err)
	}

	_ = tz
}

const dataDir = "data"

const (
	dbFileName = "TimeZoneDB.csv.zip"
	dbURL      = "https://timezonedb.com/files/" + dbFileName
)

func prepareData() error {
	downloadPath := path.Join(os.TempDir(), dbFileName)
	if _, err := os.Stat(downloadPath); err == nil {
		if err := os.Remove(downloadPath); err != nil {
			return fmt.Errorf("remove old timezone db error: %w", err)
		}
	}

	resp, err := http.Get(dbURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(downloadPath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("create temp file error: %w", err)
	}
	defer f.Close()

	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading")
	if _, err := io.Copy(io.MultiWriter(f, bar), resp.Body); err != nil {
		return fmt.Errorf("download error: %w", err)
	}

	cmd := exec.Command("unzip", "-o", "-d", dataDir, downloadPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("unzip error: %w", err)
	}

	return nil
}

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
	timeZoneStartUTC
	timeZoneOffset
	timeZoneIsDST
)

func loadCountries() (map[string]*Country, error) {
	countries := make(map[string]*Country, 10)
	f, err := os.Open(path.Join(dataDir, countryFile))
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		countries[row[countryCode]] = &Country{
			Code: row[countryCode],
			Name: row[countryName],
		}
	}
	return countries, nil
}

func loadTimeZones(countries map[string]*Country) (map[string]*Zone, error) {
	zones := make(map[string]*Zone, 10)
	f, err := os.Open(path.Join(dataDir, timeZoneFile))
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
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
			zone = &Zone{
				Name: row[timeZoneName],
			}
			zones[row[timeZoneName]] = zone
		}

		country := countries[row[timeZoneCountryCode]]
		if ok := lo.ContainsBy(zone.Countries, func(v *Country) bool {
			return v.Code == country.Code
		}); !ok {
			zone.Countries = append(zone.Countries, country)
		}
		if ok := lo.ContainsBy(country.Zones, func(v *Zone) bool {
			return v.Name == zone.Name
		}); !ok {
			country.Zones = append(country.Zones, zone)
		}

		zoneTran := &ZoneTran{
			Abbreviation: row[timeZoneAbbreviation],
			StartUTC:     cast.ToInt64(row[timeZoneStartUTC]),
			Offset:       cast.ToInt(row[timeZoneOffset]),
			IsDST:        row[timeZoneIsDST] == "1",
		}
		zone.Trans = append(zone.Trans, zoneTran)
	}

	return zones, nil
}
