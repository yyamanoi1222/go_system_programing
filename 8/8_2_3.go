package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "socket-sample")
	conn, err := net.Dial("unix", path)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:8888", nil)

	if err != nil {
		panic(err)
	}

	request.Write(conn)

	response, err := http.ReadResponse(
		bufio.NewReader(conn), request)

	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))
}
