package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	gzw := gzip.NewWriter(w)
	writer := io.MultiWriter(os.Stdout, gzw)

	encoder := json.NewEncoder(writer)
	encoder.Encode(source)

	gzw.Flush()

	// io.WriteString(w, "ok test")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
