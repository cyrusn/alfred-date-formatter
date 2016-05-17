package xmlBuilder

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Items struct {
	XMLName xml.Name `xml:"items"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Arg     string   `xml:"arg,attr"`
	Title   string   `xml:"title"`
}

//  CreateXML convert given time.Time value to xml format for alfred script filter workflow
func CreateXML(layouts []string, result time.Time) (string, error) {
	items := []Item{}
	for _, l := range layouts {
		r := result.Format(l)
		v := Item{
			Arg:   r,
			Title: r,
		}
		items = append(items, v)
	}
	f := Items{Item: items}

	output, err := xml.MarshalIndent(f, "", "")
	o := fmt.Sprintf("%s%s\n", xml.Header, output)
	return o, err
}
