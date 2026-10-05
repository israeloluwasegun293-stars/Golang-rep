package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello, welcome to this webpage, this is the homepage"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display the snippet with the id number...%d", id)

}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/create" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Write([]byte("We create snippet here....."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet", showSnippet)

	log.Println("Starting server at port :4040....")
	err := http.ListenAndServe(":4040", mux)
	log.Fatal(err)

}
