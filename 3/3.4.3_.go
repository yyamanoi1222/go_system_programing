package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHOST: ascii.jp\r\n\r\n"))

	res, err := http.ReadResponse(bufio.NewReader(conn), nil)

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
