package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
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
	topic    string
	Location []string
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

func newAgg(w http.ResponseWriter, request *http.Request) {
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
			x.Location = append(x.Location, n)
		}
	}

	t, _ := template.ParseFiles("basichtml.html")
	fmt.Println(t.Execute(w, x))
}

func index_handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "<h1>Namastey _/\\_</h1>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/News", newAgg)
	http.ListenAndServe(":8000", nil)

}
