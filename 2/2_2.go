package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("sample.txt")
	writer := io.MultiWriter(os.Stdout, f)
	w := csv.NewWriter(writer)
	w.Write([]string{"1", "2", "3"})
	w.Flush()
}
