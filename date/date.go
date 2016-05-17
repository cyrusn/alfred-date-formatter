package date

import (
	"fmt"
	"regexp"

	"strings"
	"time"
)

var regexpStrings = []string{
	`^(?P<year>\d{4})[\.|\-|\s|\/](?P<month>\d{1,2})[\.|\-|\s|\/](?P<day>\d{1,2})$`,
	`^(?P<day>\d{1,2})[\.|\-|\s|\/](?P<month>\d{1,2})[\.|\-|\s|\/]*(?P<year>\d{4}|\d{2}|\d{0})$`,
	`^(?P<day>\d{1,2})[\.|\-|\s|\/](?P<month>\w+)[\.|\-|\s|\/]*(?P<year>\d{4}|\d{2}|\d{0})$`,
}

// indexOf can search a string in []string, return index or -1 if no match found
func indexOf(slice []string, v string) int {
	for i, s := range slice {
		if s == strings.ToLower(v) {
			return i
		}
	}
	return -1
}

// ParseDate parse the date string and return time.Time format, will return bool if any unmatched case
func ParseDateString(s string) time.Time {
	for _, r := range regexpStrings {
		t, b := parseDateByGivenRegExp(r, s)
		if b {
			return t
		}
	}
	return time.Now()
}

// parseDateByGivenRegExp parse the date string with given regexp
func parseDateByGivenRegExp(reg string, date string) (time.Time, bool) {
	r := regexp.MustCompile(reg)

	if r.MatchString(date) {
		match := r.FindStringSubmatch(date)

		var result = make(map[string]string)
		for i, name := range r.SubexpNames() {
			result[name] = match[i]
		}

		day := result["day"]
		month := result["month"]
		year := result["year"]

		var layout string

		layout = "2"

		// parse month string
		switch {
		case len(month) < 3:
			layout += "-1"
		case len(month) == 3:
			layout += "-Jan"
		case len(month) > 3:
			layout += "-January"
		}

		// parse year string
		switch {
		case len(year) == 0:
			layout += "-2006"
			year = fmt.Sprintf("%v", time.Now().Year())
		case len(year) == 2:
			layout += "-06"
		case len(year) == 4:
			layout += "-2006"
		}

		value := fmt.Sprintf("%v-%v-%v", day, month, year)
		t, _ := time.Parse(layout, value)

		if t.Before(time.Now()) && len(result["year"]) == 0 {
			t = t.AddDate(1, 0, 0)
		}

		return t, true
	}

	return time.Now(), false
}
