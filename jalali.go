package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

// TimeToJalali convert time to jalali date
func TimeToJalali(t time.Time) (ptime.Time, error) {
	return ptime.Unix(t.Unix(), t.UnixNano()), nil
}

// ParseJalali parse jalali date from string
func ParseJalali(str string, loc *time.Location) *ptime.Time {
	isNumber := func(v string) bool { return regexp.MustCompile(`^[0-9]+$`).MatchString(v) }

	pattern := regexp.MustCompile(`[^\d]+`)
	_parts := strings.Split(pattern.ReplaceAllString(str, "-"), "-")
	parts := make([]int, 6)
	for i := 0; i < 6; i++ {
		item := "0"
		if len(_parts) > i && isNumber(_parts[i]) {
			item = _parts[i]
		}
		t, _ := strconv.Atoi(item)
		parts[i] = t
	}

	if loc == nil {
		loc = time.FixedZone("Asia/Tehran", 12600)
	}

	date := ptime.Date(parts[0], ptime.Month(parts[1]), parts[2], parts[3], parts[4], parts[5], 0, loc)
	if date.Year() == parts[0] && int(date.Month()) == parts[1] && date.Day() == parts[2] {
		return &date
	}
	return nil
}
