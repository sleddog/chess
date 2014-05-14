package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
)

func randomMove(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Content-type: text/json\n\n")
	fmt.Fprintf(w, "{")
	fmt.Fprintf(w, "\"query\":\"blahhh\",")
	fmt.Fprintf(w, getNextMove())
	fmt.Fprintf(w, "}\n")
}

func main() {
	http.HandleFunc("/", randomMove)
	cgi.Serve(nil)
}
