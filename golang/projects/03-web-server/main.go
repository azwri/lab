package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fielServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fielServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandller)

	fmt.Println("Server start at port 8080. http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "/hello")
}

func formHandller(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error %v\n", err)
	}
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Post request successfull.\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s.\nAddress: %s.\n", name, address)
}
