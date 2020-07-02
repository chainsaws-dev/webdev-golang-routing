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
		// Can you answer the question as to why "I see you connected." is never written?
		//
		// YES. BECAUSE WE NEVER LEAVE THIS LOOP IT CONTINUES TO EXPECT NEW LINES
		// WE SHOULD CHECK IF ln IS EMPTY STRING AND BREAK IF IT IS
		i++
	}

	//Never gets here because we never stop scanning
	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")

}
