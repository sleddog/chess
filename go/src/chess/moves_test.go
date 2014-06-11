package chess

import (
	"fmt"
	"testing"
)

func TestGetMoves(t *testing.T) {
	node := createChessNode(initial_board_json)
	legal_black_moves := getMoves("b", node.board)
	fmt.Println("legal_black_moves = ", legal_black_moves)
	fmt.Println("count = ", len(legal_black_moves))
	if len(legal_black_moves) != 20 {
		t.Error("Initial moves for black should be 20")
	}
}

func TestGetLegalMoves(t *testing.T) {
	node := createChessNode(initial_board_json)
	legal_black_moves := getLegalMoves("b", node.board)
	fmt.Println("legal_black_moves = ", legal_black_moves)
	fmt.Println("count = ", len(legal_black_moves))
	if len(legal_black_moves) != 20 {
		t.Error("Initial moves for black should be 20")
	}
}
