package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

// JalaliToTime convert jalali date string to time from 2006-01-02 format
func JalaliToTime(jDate string) (time.Time, error) {
	t := time.Time{}
	rx := regexp.MustCompile(`^\d{4}(-|\/)(\d{1,2})(-|\/)(\d{1,2})$`)
	if !rx.MatchString(jDate) {
		return t, errors.New("Inalid date format")
	}

	jDate = strings.ReplaceAll(jDate, "/", "-")
	jDate = strings.ReplaceAll(jDate, "\\", "-")
	jDate = strings.ReplaceAll(jDate, ".", "-")
	parts := strings.Split(jDate, "-")

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return t, errors.New("invalid year")
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil || (month > 12 || month < 1) {
		return t, errors.New("invalid month")
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil || (month > 31 || month < 1) {
		return t, errors.New("invalid day")
	}

	pt := ptime.Date(year, ptime.Month(month), day, 0, 0, 0, 0, ptime.Iran())
	if pt.Year() != year || int(pt.Month()) != month || pt.Day() != day {
		return t, errors.New("Inalid date")
	}

	return pt.Time(), nil
}

// TimeToJalali convert time to jalali date
func TimeToJalali(t time.Time) (ptime.Time, error) {
	return ptime.Unix(t.Unix(), t.UnixNano()), nil
}
