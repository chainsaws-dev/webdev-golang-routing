package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func main() {

	http.Handle("/", http.HandlerFunc(roothandler))
	http.Handle("/dog/", http.HandlerFunc(doghandler))
	http.Handle("/me/", http.HandlerFunc(mehandler))

	http.ListenAndServe(":8080", nil)
}

func roothandler(rw http.ResponseWriter, req *http.Request) {
	tmp, err := template.ParseFiles("root.html")

	if err != nil {
		log.Fatalln(err)
	}

	tmp.ExecuteTemplate(rw, "root.html", tmp)
}

func doghandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Dog says woof!")
}

func mehandler(rw http.ResponseWriter, req *http.Request) {

	tmp, err := template.ParseFiles("me.html")

	if err != nil {
		log.Fatalln(err)
	}

	au := strings.Split(req.RequestURI, "/")

	if len(au) >= 2 {
		if au[2] != "" {
			tmp.ExecuteTemplate(rw, "me.html", au[2])
		} else {
			tmp.ExecuteTemplate(rw, "me.html", "Anon")
		}
	}

}
