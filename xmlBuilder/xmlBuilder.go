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
	Uid      int      `xml:"uid,attr"`
	Arg      string   `xml:"arg,attr"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
}

//  CreateXML convert given time.Time value to xml format for alfred script filter workflow
func CreateXML(layouts []string, result time.Time) (string, error) {
	items := []Item{}
	for i, l := range layouts {
		r := result.Format(l)
		v := Item{
			Uid:      i,
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
