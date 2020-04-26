package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemap struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s Sitemap
	xml.Unmarshal(bytes, &s)

	for _, n := range s.Locations {
		fmt.Printf("%s", n)
	}
}
