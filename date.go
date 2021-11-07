package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

// JalaliToTime convert jalali date string to time from 2006-01-02 15:04:05 format
//
// this function returns nil on invalid date
func JalaliToTime(jd string) *time.Time {
	isNumber := func(v string) bool { return regexp.MustCompile(`^[0-9]+$`).MatchString(v) }
	p := func(str string) []int {
		pattern := regexp.MustCompile(`[^\d]+`)
		parts := strings.Split(pattern.ReplaceAllString(str, "-"), "-")
		res := make([]int, 6)
		for i := 0; i < 6; i++ {
			item := "0"
			if len(parts) > i && isNumber(parts[i]) {
				item = parts[i]
			}
			t, _ := strconv.Atoi(item)
			res[i] = t
		}
		return res
	}(jd)

	date := ptime.Date(p[0], ptime.Month(p[1]), p[2], p[3], p[4], p[5], 0, time.FixedZone("Asia/Tehran", 12600))
	if date.Year() == p[0] && int(date.Month()) == p[1] && date.Day() == p[2] {
		res := date.Time()
		return &res
	}
	return nil
}

// TimeToJalali convert time to jalali date
func TimeToJalali(t time.Time) (ptime.Time, error) {
	return ptime.Unix(t.Unix(), t.UnixNano()), nil
}
