package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("get", "http://localhost:8888", nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, _ := httputil.DumpResponse(response, true)
	fmt.Println(string(dump))
}
