package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	sm := http.NewServeMux()

	sm.HandleFunc("/", roothandler)
	sm.HandleFunc("/dog/", doghandler)
	sm.HandleFunc("/me/", mehandler)

	http.ListenAndServe(":8080", sm)
}

func roothandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, `<p>Server is running</p><br><a href="/dog">dog</a> <a href="/me">me</a>`)
}

func doghandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Dog says woof!")
}

func mehandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, req.UserAgent())

	au := strings.Split(req.RequestURI, "/")

	if len(au) >= 2 {
		if au[2] != "" {
			fmt.Fprintln(rw, "Hello,", strings.ToUpper(au[2]), "!")
		} else {
			fmt.Fprintln(rw, "Hello, Anon!")
		}
	}

}
