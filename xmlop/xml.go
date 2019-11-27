package main

import (
	"encoding/xml"
	"fmt"
)

type Struct struct {
	XMLName   xml.Name `xml:"struct"`
	Name      string   `xml:"userShare,attr"`
	EventType string   `xml:"eventtype,attr"`
	Group     string   `xml:"group,attr"`
	Entries   []Entry  `xml:"entry"`
}

type Entry struct {
	Name  string `xml:"name,attr"`
	Type  string `xml:"type,attr"`
	Desc  string `xml:"desc,attr"`
	Value int    `xml:",chardata"`
}

func main() {
	e1 := Entry{Name: "share_channel", Type: "int", Desc: "分享渠道", Value: 123}
	e2 := Entry{Name: "share_type", Type: "int", Desc: "分享类型", Value: 345}
	s := &Struct{
		Name:      "userShare",
		EventType: "eventtype",
		Group:     "CommonField",
		Entries:   []Entry{e1, e2},
	}

	output, err := xml.MarshalIndent(s, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to xml: ", err)
		return
	}
	fmt.Println(string(output))

}
