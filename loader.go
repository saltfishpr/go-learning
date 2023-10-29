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

	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cast"
)

type loader struct {
	downloadData bool
	baseURL      string // https://timezonedb.com
	fileName     string // TimeZoneDB.csv.zip
}

func (l *loader) Load() ([]*Zone, error) {
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

	res := make([]*Zone, 0, len(zones))
	for _, key := range keys {
		res = append(res, zones[key])
	}
	return res, nil
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

func (l *loader) load(dataDir string) (map[string]*Zone, error) {
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
