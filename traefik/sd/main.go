package main

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed config.yml
var config string

func loadConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(config))
}

func main() {
	http.HandleFunc("/api", loadConfig)
	http.ListenAndServe(":9000", nil)
}
