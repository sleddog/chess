package main

import (
	"chess"
	"fmt"
	"net/http"
	"net/http/cgi"
	//  "reflect"
)

func randomMove(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Content-type: text/json\n\n")
	fmt.Fprintf(w, "{")
	query := r.URL.Query()
	board := query["board"]
	//fmt.Fprintf(w, "\"query\":\"", query, "\",")
	//fmt.Fprintf(w, "\"board\":\"",board,"\",")
	//fmt.Fprintf(w, chess.GetNextMoveUsingArray(board))
	//fmt.Println("type:", reflect.TypeOf(board))
	//fmt.Println(board)
	//fmt.Printf("%v", board)
	//for i:=0; i<len(board); i++ {
	//        foo := board[i]
	//        fmt.Println("i=",i,", board[i]=",foo)
	//        fmt.Println("foo type:", reflect.TypeOf(foo))
	//}
	//fmt.Fprintf(w, chess.GetNextMoveUsingJson(board[0]))
	//TODO pass in the board properly... currently a []string, needed to be just a string
	//fmt.Fprintf(w, chess.GetNextMove())

	fmt.Fprintf(w, chess.GetNextMoveUsingPointValue(board))

	fmt.Fprintf(w, "}\n")
}

func main() {
	http.HandleFunc("/", randomMove)
	cgi.Serve(nil)
}
