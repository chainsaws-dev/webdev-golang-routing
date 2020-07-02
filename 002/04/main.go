package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()

		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	sc := bufio.NewScanner(conn)

	i := 1

	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(i, ln)

		if ln == "" {
			break
		}

		i++
	}

	//Never gets here because we never stop scanning
	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")

	// WE CANNOT WRITE STRINGS WITHOUT FOLOWING THE RULES OF RESPONSE MARKUP OTHERWISE NOTHING WILL BE DISPLAYED IN BROWSER

}
