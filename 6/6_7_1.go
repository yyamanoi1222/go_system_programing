package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {

	current := 0
	var conn net.Conn = nil
	sendMessages := []string{
		"A",
		"B",
		"C",
	}

	for {
		var err error

		if current == 1 {
			time.Sleep(time.Second * 6)
		}

		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")

			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		request, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		request.Header.Set("Accept-Encoding", "gzip")
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)

		if err != nil {
			panic(err)
		}

		response, err := http.ReadResponse(
			bufio.NewReader(conn), request)

		if err != nil {
			fmt.Println("retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(response, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer response.Body.Close()

		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, response.Body)
		}

		current++

		if current == len(sendMessages) {
			break
		}
	}

	conn.Close()
}
