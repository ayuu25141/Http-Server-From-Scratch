package main

import (
	"fmt"
	"net"
	"strings"
)

// Handler type
type HandlerFunc func() string

// Router
var routes = map[string]HandlerFunc{
	"/":        homeHandler,
	"/health":  healthHandler,
	"/about":   aboutHandler,
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on http://localhost:8080")

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	req := string(buf[:n])

	// Parse request line
	lines := strings.Split(req, "\r\n")
	parts := strings.Split(lines[0], " ")

	method := parts[0]
	path := parts[1]

	handler, exists := routes[path]

	body := ""
	status := "200 OK"

	if method != "GET" {
		status = "405 Method Not Allowed"
		body = "Method Not Allowed"
	} else if exists {
		body = handler()
	} else {
		status = "404 Not Found"
		body = "Not Found"
	}

	response := fmt.Sprintf(
		"HTTP/1.1 %s\r\nContent-Length: %d\r\n\r\n%s",
		status,
		len(body),
		body,
	)

	conn.Write([]byte(response))
}
func homeHandler()  string {
	return  "Homepage"
}
func healthHandler()  string {
	return  "Ok Health"
}
func aboutHandler()  string {
	return  "About page"
}
