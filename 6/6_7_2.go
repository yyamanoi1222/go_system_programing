package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func isGzipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
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

		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header:     make(http.Header),
		}

		if isGzipAcceptable(request) {
			content := "herro world (gzipped)"
			var buffer bytes.Buffer

			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			response.Body = ioutil.NopCloser(&buffer)
			response.ContentLength = int64(buffer.Len())
			response.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "herro world"
			response.Body = ioutil.NopCloser(strings.NewReader(content))
			response.ContentLength = int64(len(content))
		}

		response.Write(conn)
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
