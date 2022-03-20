package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
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

		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		current++

		if current == len(sendMessages) {
			break
		}
	}

	conn.Close()
}
