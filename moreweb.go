package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<h1>Namastey samajwasio _/\_</h1>
					<p>Go is Fast !!!</p>
					<p> .... and simple!</p>`)

	/*fmt.Fprintf(w, "<h1>Namastey samajwasio _/\\_</h1>")
	fmt.Fprintf(w, "<p>Go is Fast !!!</p>")
	fmt.Fprintf(w, "<p> .... and simple!</p>")*/

}

func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
