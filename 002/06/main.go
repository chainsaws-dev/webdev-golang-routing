package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn) {

	defer conn.Close()

	request(conn)

	response(conn)
}

func request(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		log.Println(err)
	}

	newscanner := bufio.NewScanner(conn)
	i := 0

	for newscanner.Scan() {
		ln := newscanner.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
		i++
	}

}

func response(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="ru"><head><meta charset="UTF-8"><title>Состояние сервера</title></head><body><strong>Сервер работает!</strong></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
