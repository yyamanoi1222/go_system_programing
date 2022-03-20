package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

var contents = []string{
	"aaaaa",
	"bbbbb",
	"ccccc",
}

func processSession(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Accept %v \n", conn.RemoteAddr())

	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		request, err := http.ReadRequest(bufio.NewReader(conn))

		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}

		dump, err := httputil.DumpRequest(request, true)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))

		for _, content := range contents {
			bytes := []byte(content)
			time.Sleep(2 * time.Second)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}

		fmt.Fprintf(conn, "0\r\n\r\n")
	}

}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go processSession(conn)
	}
}
