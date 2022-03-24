package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("test test test")

	sectionReader := io.NewSectionReader(reader, 3, 6)

	io.Copy(os.Stdout, sectionReader)
}
