package main

import (
	"./date"
	"encoding/json"
	"flag"
	"fmt"
	"time"
)

// var day_change int
var flagvar string

func main() {
	flag.StringVar(&flagvar, "d", time.Now().Format("2006.01.02"), "\"date string\"")
	flag.Parse()

	output, _ := jsonFormatter(flagvar)

	fmt.Printf("%s\n", output)
}

func jsonFormatter(s string) ([]byte, error) {

	type Item struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}

	type Items struct {
		Items []Item `json:"items"`
	}

	var layouts = []string{
		"2006.01.02 (Mon)",
		"2006-01-02 (Mon)",
		"02.01.06 (Mon)",
		"02/01/06 (Mon)",
		"2 Jan 06 (Mon)",
		"Mon, 2 Jan 06",
	}

	result := Items{
		Items: []Item{},
	}

	for i, l := range layouts {
		t := Item{
			Id:    i,
			Title: date.ParseDateString(s).Format(l),
		}
		result.Items = append(result.Items, t)
	}

	return json.Marshal(result)
}
