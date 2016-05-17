package date

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var monthNames = []string{
	"jan",
	"feb",
	"mar",
	"apr",
	"may",
	"jun",
	"jul",
	"aug",
	"sep",
	"oct",
	"nov",
	"dec",
	"january",
	"february",
	"march",
	"april",
	"may",
	"june",
	"july",
	"august",
	"september",
	"october",
	"november",
	"december",
}

func indexOf(slice []string, v string) int {
	for i, s := range slice {
		if s == strings.ToLower(v) {
			return i
		}
	}
	return -1
}

func ParseDate(reg string, date string) (time.Time, bool) {
	r := regexp.MustCompile(reg)

	if r.MatchString(date) {
		match := r.FindStringSubmatch(date)

		var result = make(map[string]string)
		for i, name := range r.SubexpNames() {
			result[name] = match[i]
		}

		year, _ := strconv.Atoi(result["year"])
		day, _ := strconv.Atoi(result["day"])
		month, _ := strconv.Atoi(result["month"])

		// manage year
		now := time.Now()
		if year == 0 {
			if int(now.Month()) <= month && now.Day() <= day {
				year = now.Year()
			} else {
				year = now.Year() + 1
			}
		} else if year < 100 {
			year = year + 2000
		}

		if month == 0 {
			position := indexOf(monthNames, result["month"])
			month = (position % 12) + 1
		}

		layout := "2006-1-2"

		value := fmt.Sprintf("%v-%v-%v", year, month, day)
		t, _ := time.Parse(layout, value)
		return t, true
	}

	return time.Now(), false
}
