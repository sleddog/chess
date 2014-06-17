package chess

import (
	"fmt"
	"testing"
)

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
	node.legal_moves = getMoves("b", node.board)

	board2 := makeMove(node.board, node.legal_moves[0])
	fmt.Println("board2=", board2)
}

func TestMiniMax(t *testing.T) {
	fenPlacement := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)

	node := ChessNode{board: board, depth: 0, active_color: "w"}
	move := miniMaxDecision(node)
	fmt.Println("move = ", move)
}

func TestMiniMax2(t *testing.T) {
	fenPlacement := "rn1qkbnr/p1p1p1pp/3p1p2/1p1P4/8/2NB1NP1/PPPP1P1P/R1BQ1RK1"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)

	node := ChessNode{board: board, depth: 0, active_color: "b"}
	move := miniMaxDecision(node)
	fmt.Println("move = ", move)
}

func TestMiniMax3(t *testing.T) {
	fenPlacement := "rnbqkbnr/p1pppppp/8/1p6/2B1P3/8/PPPP1PPP/RNBQK1NR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)

	node := ChessNode{board: board, depth: 0, active_color: "b"}
	move := miniMaxDecision(node)
	fmt.Println("move = ", move)
}
