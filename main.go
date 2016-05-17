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
	output, _ := xmlBuilder.CreateXML(layouts, date.ParseDateString(flagvar))

	fmt.Printf("%s\n", output)
}

var layouts = []string{
	"2006.01.02 (Mon)",
	"2006-01-02 (Mon)",
	"02.01.06 (Mon)",
	"02/01/06 (Mon)",
	"2 Jan 06 (Mon)",
	"Mon, 1 Jan 06",
}
