package internal

import (
	"time"
)

func TimeString2Timestamp(s string, offset int) (int64, error) {
	t, err := time.ParseInLocation(time.DateTime, s, time.UTC)
	if err != nil {
		return 0, err
	}
	return t.Add(time.Duration(-offset) * time.Minute).Unix(), nil
}

func TimestampToTimeString(sec int64, offset int) string {
	return time.Unix(sec, 0).UTC().Add(time.Duration(offset) * time.Minute).Format(time.DateTime)
}
