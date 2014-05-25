package chess

import (
	"fmt"
	"testing"
)

func TestGetNextMove(t *testing.T) {
	move := GetNextMove()
	fmt.Println(move)
}

func TestCreateBoard(t *testing.T) {
	board := createBoard(initial_board_json)
	fmt.Println(board)
}

func TestCreateChessNode(t *testing.T) {
	node := createChessNode(initial_board_json)
	fmt.Println("node = ", node)
}

func TestPrintNode(t *testing.T) {
	node := createChessNode(initial_board_json)
	printNode(node)
}

func TestGetLegalBlackMoves(t *testing.T) {
	node := createChessNode(initial_board_json)
	legal_black_moves := getLegalBlackMoves(node)
	fmt.Println("legal_black_moves = ", legal_black_moves)
}

func TestGetNextMoveUsingArray(t *testing.T) {
	var dat []string
	dat = append(dat, "[[\"wr\",\"wn\",\"wb\",\"wq\",\"wk\",\"wb\",\"wn\",\"wr\"],[\"wp\",\"wp\",\"wp\",\"wp\",0,\"wp\",\"wp\",\"wp\"],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[\"bp\",0,0,0,\"wp\",0,0,0],[0,\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\"],[\"br\",\"bn\",\"bb\",\"bq\",\"bk\",\"bb\",\"bn\",\"br\"]]")
	fmt.Println("dat=", dat)
	move := GetNextMoveUsingArray(dat)
	fmt.Println("move=", move)
}
