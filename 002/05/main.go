package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	m, curi := request(conn)

	distext := fmt.Sprintln(m, curi)

	response(conn, distext)
}

func request(conn net.Conn) (method string, uri string) {
	var m string
	var resuri string
	var p string

	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		log.Println(err)
	}

	newscanner := bufio.NewScanner(conn)
	i := 0

	for newscanner.Scan() {
		ln := newscanner.Text()
		//fmt.Println(ln)
		if i == 0 {
			fld1 := strings.Fields(ln)
			m = fld1[0]
			p = fld1[1]
		}

		if i == 1 {
			fld2 := strings.Fields(ln)
			resuri = "http://" + fld2[1]
			resuri += p
		}

		if ln == "" {
			break
		}
		i++
	}

	fmt.Println(m, resuri)

	return m, resuri
}

func response(conn net.Conn, displaytext string) {

	body := `<!DOCTYPE html><html lang="ru"><head><meta charset="UTF-8"><title>Состояние сервера</title></head><body><strong>Сервер работает!</strong></body></html>`
	body = strings.Replace(body, `<strong>Сервер работает!</strong>`, `<strong>Сервер работает!</strong><ul><li>`+displaytext+`</li></ul>`, 1)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
