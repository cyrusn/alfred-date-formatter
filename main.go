package main

import (
	"./date"
	"./xmlBuilder"

	// "errors"
	"flag"
	"fmt"
	// "regexp"
	"time"
)

// var day_change int
var flagvar string

func main() {
	flag.StringVar(&flagvar, "d", time.Now().Format("2006.01.02"), "help message for flagname")

	flag.Parse()

	// result, _ := date.ParseDate(flagvar)
	output, _ := xmlBuilder.CreateXML(layouts, parseArgv(flagvar))

	fmt.Printf("%s\n", output)
}

var layouts = []string{
	"2006.1.2 (Mon)",
	"2 Jan 2006 (Mon)",
}

var regexpStrings = []string{
	`^(?P<year>\d{4})[\.|\-|\s|\/](?P<month>\d{1,2})[\.|\-|\s|\/](?P<day>\d{1,2})$`,
	`^(?P<day>\d{1,2})[\.|\-|\s|\/](?P<month>\d{1,2})[\.|\-|\s|\/]*(?P<year>\d{4}|\d{2}|\d{0})$`,
	`^(?P<day>\d{1,2})[\.|\-|\s|\/](?P<month>\w+)[\.|\-|\s|\/]*(?P<year>\d{4}|\d{2}|\d{0})$`,
}

func parseArgv(argv string) time.Time {
	for _, r := range regexpStrings {
		t, b := date.ParseDate(r, argv)
		if b {
			return t
		}
	}
	return time.Now()
}

// yearComplement will add year(future) to t if t have have no year value
func yearComplement(t time.Time) time.Time {
	var r time.Time

	if t.Year() == 0 {
		if t.AddDate(time.Now().Year(), 0, 0).Before(time.Now()) {
			r = t.AddDate(time.Now().Year()+1, 0, 0)
		} else {
			r = t.AddDate(time.Now().Year(), 0, 0)
		}
	} else {
		r = t
	}
	return r
}
