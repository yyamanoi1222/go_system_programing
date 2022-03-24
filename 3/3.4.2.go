package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("3.4.1.go")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	io.Copy(os.Stdout, file)
}
