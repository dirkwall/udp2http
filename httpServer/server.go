package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	fmt.Fprintf(w, "200 OK\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":51000", nil)
}

