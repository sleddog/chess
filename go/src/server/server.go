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
//	fmt.Fprintf(w, "\"query\":\"", r.RequestURI, "\",")
	fmt.Fprintf(w, "\"query\":\"ohhhyeaaahh\",")
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
