package chess

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//define global variables
var pieces = make(map[string]map[string]string)
var initial_board_json string

func init() {
	fmt.Println("Inside init()")

	//set the random seed
	rand.Seed(time.Now().UTC().UnixNano())

	//initialize the html and codepoint value for each piece
	pieces = map[string]map[string]string{
		"wk": {"html": "&#9812;", "codepoint": "\u2654"},
		"wq": {"html": "&#9813;", "codepoint": "\u2655"},
		"wr": {"html": "&#9814;", "codepoint": "\u2656"},
		"wb": {"html": "&#9815;", "codepoint": "\u2657"},
		"wn": {"html": "&#9816;", "codepoint": "\u2658"},
		"wp": {"html": "&#9817;", "codepoint": "\u2659"},
		"bk": {"html": "&#9818;", "codepoint": "\u265A"},
		"bq": {"html": "&#9819;", "codepoint": "\u265B"},
		"br": {"html": "&#9820;", "codepoint": "\u265C"},
		"bb": {"html": "&#9821;", "codepoint": "\u265D"},
		"bn": {"html": "&#9822;", "codepoint": "\u265E"},
		"bp": {"html": "&#9823;", "codepoint": "\u265F"}}
	fmt.Println("pieces = ", pieces)

	//set the initial_board_json
	initial_board_json = "{\"a8\":\"br\",\"b8\":\"bn\",\"c8\":\"bb\",\"d8\":\"bq\",\"e8\":\"bk\",\"f8\":\"bb\",\"g8\":\"bn\",\"h8\":\"br\",\"a7\":\"bp\",\"b7\":\"bp\",\"c7\":\"bp\",\"d7\":\"bp\",\"e7\":\"bp\",\"f7\":\"bp\",\"g7\":\"bp\",\"h7\":\"bp\",\"a2\":\"wp\",\"b2\":\"wp\",\"c2\":\"wp\",\"d2\":\"wp\",\"e2\":\"wp\",\"f2\":\"wp\",\"g2\":\"wp\",\"h2\":\"wp\",\"a1\":\"wr\",\"b1\":\"wn\",\"c1\":\"wb\",\"d1\":\"wq\",\"e1\":\"wk\",\"f1\":\"wb\",\"g1\":\"wn\",\"h1\":\"wr\"}"
}

//struct to hold all data about a particular board
type ChessNode struct {
	board             [8][8]string
	white_legal_moves []string
	black_legal_moves []string
}

func createChessNode(board_json string) ChessNode {
	node := ChessNode{board: createBoard(board_json)}
	return node
}

func randomColumn() string {
	var columns string
	columns = "abcdefgh"
	return string(columns[rand.Intn(len(columns))])
}

func getNextMove() string {
	var move string
	var randChar string
	randChar = randomColumn()
	move = fmt.Sprintf("\"next-move\":\"%s7-%s5\"", randChar, randChar)
	return move
}

//return an 8x8 board from the JSON representation
func createBoard(board_json string) [8][8]string {
	//convert json to string map
	byt := []byte(board_json)
	var dat map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	//populate board matrix
	var board [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			letter := numberToLetter(j)
			square := letter + strconv.Itoa(i+1)
			if val, ok := dat[square]; ok {
				board[i][j] = val
			} else {
				board[i][j] = "0"
			}
		}
	}
	return board
}

func numberToLetter(x int) string {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return letters[x]
}

func printNode(node ChessNode) {
	board := node.board
	for i := 7; i >= 0; i-- {
		row := ""
		for j := 0; j < 8; j++ {
			if board[i][j] == "0" {
				row += "\u3000"
			} else {
				row += pieceToUnicode(board[i][j])
			}
		}
		fmt.Println(row)
	}
}

func pieceToUnicode(piece string) string {
	return pieces[piece]["codepoint"]
}

//for the given ChessNode return all of the legal black moves
func getLegalBlackMoves(node ChessNode) []string {
	var black_moves []string
	board := node.board
	for i := 7; i >= 0; i-- {
		for j := 0; j < 8; j++ {
			piece := board[i][j]
			if piece == "0" {
				continue
			} else if piece[0:1] == "b" {
				//found a black piece
				fmt.Println("black piece = ", piece, ", i=", i, "j=", j)
				black_moves = append(black_moves, getMovesForBlackPiece(piece, i, j, node)...)
			}
		}
	}
	fmt.Println("# of legal black_moves = ", len(black_moves))
	return black_moves
}

func getMovesForBlackPiece(piece string, i int, j int, node ChessNode) []string {
	var moves []string
	piece_type := piece[1:2]
	switch piece_type {
	case "p":
		fmt.Println("PAWN")
		moves = append(moves, getMovesForBlackPawn(i, j, node)...)
	default:
		fmt.Println("DEFAULT = ", piece)
	}
	return moves
}

func getMovesForBlackPawn(i int, j int, node ChessNode) []string {
	var moves []string
	//TODO check if first square in front is blank
	if i > 0 && node.board[i-1][j] == "0" {
		moves = append(moves, coordToSquare(i, j))
	}
	return moves
}

func coordToSquare(i int, j int) string {
	return "TODO"
}
