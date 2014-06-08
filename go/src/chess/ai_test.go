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
	fmt.Println("count = ", len(legal_black_moves))
}

func TestGetNextMoveUsingArray(t *testing.T) {
	var dat []string
	dat = append(dat, "[[\"wr\",\"wn\",\"wb\",\"wq\",\"wk\",\"wb\",\"wn\",\"wr\"],[\"wp\",\"wp\",\"wp\",\"wp\",0,\"wp\",\"wp\",\"wp\"],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[\"bp\",0,0,0,\"wp\",0,0,0],[0,\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\"],[\"br\",\"bn\",\"bb\",\"bq\",\"bk\",\"bb\",\"bn\",\"br\"]]")
	fmt.Println("dat=", dat)
	move := GetNextMoveUsingArray(dat)
	fmt.Println("move=", move)
}

func TestParseFen(t *testing.T) {
	//fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	//TODO split on space above, for now just use placement directly
	fenPlacement := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)
	fmt.Println("board = ", board)
}

func TestGetNextMoveUsingPointValue(t *testing.T) {
	var dat []string
	dat = append(dat, "[[\"wr\",\"wn\",\"wb\",\"wq\",\"wk\",\"wb\",\"wn\",\"wr\"],[\"wp\",\"wp\",\"wp\",\"wp\",0,\"wp\",\"wp\",\"wp\"],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[0,0,0,0,0,0,0,0],[\"bp\",0,0,0,\"wp\",0,0,0],[0,\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\",\"bp\"],[\"br\",\"bn\",\"bb\",\"bq\",\"bk\",\"bb\",\"bn\",\"br\"]]")
	fmt.Println("dat=", dat)
	move := GetNextMoveUsingPointValue(dat)
	fmt.Println("move=", move)
}

func TestMakeMove(t *testing.T) {
	fenPlacement := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)
	node := ChessNode{board: board}
	node.black_legal_moves = getLegalBlackMoves(node)

	node2 := makeMove(node, node.black_legal_moves[0])
	fmt.Println("node2=", node2)
}
