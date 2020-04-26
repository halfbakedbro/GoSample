package main

import (
	"fmt"
	"net/http"
)

func index_handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Namastey _/\\_")
}

func main() {
	http.HandleFunc("/", index_handler)

	http.ListenAndServe(":8000", nil)
}
