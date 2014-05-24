package main

import (
	"chess"
	"fmt"
	"net/http"
	"net/http/cgi"
)

func randomMove(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Content-type: text/json\n\n")
	fmt.Fprintf(w, "{")
	//query := r.URL.Query()
	//board := query["board"]
	//fmt.Fprintf(w, "\"query\":\"", query, "\",")
	fmt.Fprintf(w, "\"query\":\"doit\",")
	//fmt.Fprintf(w, chess.GetNextMove(board))
	//TODO pass in the board properly... currently a []string, needed to be just a string
	fmt.Fprintf(w, chess.GetNextMove())
	fmt.Fprintf(w, "}\n")
}

func main() {
	http.HandleFunc("/", randomMove)
	cgi.Serve(nil)
}

func getNextMove() string {
	fmt.Println("getNextMove")
	return chess.GetNextMove()
}
