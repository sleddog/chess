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
