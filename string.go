package utils

import (
	"math/rand"
	"regexp"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// ExtractNumbers extract numbers from string
func ExtractNumbers(str string) string {
	rx := regexp.MustCompile("[0-9]+")
	return strings.Join(rx.FindAllString(str, -1), "")
}

// RandomStringFromCharset generate random string from character list
func RandomStringFromCharset(n uint, letters string) (res string, err error) {
	bytes := make([]byte, n)
	_, err = rand.Read(bytes)
	if err != nil {
		return
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	res = string(bytes)
	return
}

// RandomString generate random string from Alpha-Num Chars
func RandomString(n uint) (string, error) {
	return RandomStringFromCharset(n, "ABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")
}

// Slugify make slugify string
func Slugify(str ...string) string {
	r := regexp.MustCompile("\\s+")
	r2 := regexp.MustCompile("\\-+")
	return string(r2.ReplaceAllString(string(r.ReplaceAllString(strings.Join(str, "-"), "-")), "-"))
}

// ConcatStr join strings with separator
func ConcatStr(sep string, str ...string) string {
	res := make([]string, 0)
	for _, v := range str {
		if strings.TrimSpace(v) != "" {
			res = append(res, v)
		}
	}
	return strings.Join(res, sep)
}

// FormatNumber format number with comma separator
func FormatNumber(format string, v ...interface{}) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf(format, v...)
}
