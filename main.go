package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "go=%s\n", runtime.Version())
	})
	http.ListenAndServe(":8080", nil)
}
