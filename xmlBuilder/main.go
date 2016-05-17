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
	XMLName  xml.Name `xml:"item"`
	Uid      string   `xml:"uid,attr"`
	Arg      string   `xml:"arg,attr"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
}

func CreateXML(layouts []string, result time.Time) (string, error) {
	items := []Item{}
	for _, l := range layouts {
		r := result.Format(l)
		v := Item{
			Uid:      r,
			Arg:      r,
			Title:    r,
			Subtitle: r,
		}
		items = append(items, v)
	}
	f := Items{Item: items}

	output, err := xml.MarshalIndent(f, "", "")
	o := fmt.Sprintf("%s%s\n", xml.Header, output)
	return o, err
}
