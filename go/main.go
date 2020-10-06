package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", response)
	http.ListenAndServe(":8090", nil)
}

func response(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(s string) string {
	return fmt.Sprintf("<b>%s</b> a", s)
}
