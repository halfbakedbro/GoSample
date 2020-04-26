package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type Sitemap struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Loc []string `xml:"url>loc"`
}

type Data struct {
	topic string
	loc   []string
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	c <- n
}

func main() {
	var s Sitemap
	var x Data
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()

	queue := make(chan News, 100)
	for _, loc := range s.Locations {
		nn := strings.TrimSpace(loc)
		/*ss := strings.FieldsFunc(nn, func(r rune) bool {
			if r == '/' || r == '.' {
				return true
			}
			return false
		})*/
		//x.topic = ss[len(ss)-2]
		wg.Add(1)
		go newsRoutine(queue, nn)
	}

	wg.Wait()
	close(queue)

	for item := range queue {
		for _, n := range item.Loc {
			x.loc = append(x.loc, n)
		}
	}

	for _, itemitem := range x.loc {
		fmt.Println(itemitem)
	}

}
