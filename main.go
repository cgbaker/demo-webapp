package main

import (
	"fmt"
	"net/http"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"<html><h3>Version: %s</h3></html>", version())
}

func main() {
	http.HandleFunc("/", versionHandler)
	http.ListenAndServe(":8080", nil)
}
