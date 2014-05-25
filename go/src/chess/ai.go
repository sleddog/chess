package chess

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//define global variables
var pieces = make(map[string]map[string]string)
var initial_board_json string

func init() {
	//fmt.Println("Inside init()")

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
	//fmt.Println("pieces = ", pieces)

	//set the initial_board_json
	initial_board_json = "{\"a8\":\"br\",\"b8\":\"bn\",\"c8\":\"bb\",\"d8\":\"bq\",\"e8\":\"bk\",\"f8\":\"bb\",\"g8\":\"bn\",\"h8\":\"br\",\"a7\":\"bp\",\"b7\":\"bp\",\"c7\":\"bp\",\"d7\":\"bp\",\"e7\":\"bp\",\"f7\":\"bp\",\"g7\":\"bp\",\"h7\":\"bp\",\"a2\":\"wp\",\"b2\":\"wp\",\"c2\":\"wp\",\"d2\":\"wp\",\"e2\":\"wp\",\"f2\":\"wp\",\"g2\":\"wp\",\"h2\":\"wp\",\"a1\":\"wr\",\"b1\":\"wn\",\"c1\":\"wb\",\"d1\":\"wq\",\"e1\":\"wk\",\"f1\":\"wb\",\"g1\":\"wn\",\"h1\":\"wr\"}"
}

//struct to hold all data about a particular board
type ChessNode struct {
	board             [8][8]string
	white_legal_moves []Move
	black_legal_moves []Move
}

type Move struct {
	orig_row int
	orig_col int
	dest_row int
	dest_col int
	piece    string
}

func formatNextMove(move Move) string {
	columns := "abcdefgh"
	fromSquare := string(columns[move.orig_col]) + strconv.Itoa(move.orig_row+1)
	toSquare := string(columns[move.dest_col]) + strconv.Itoa(move.dest_row+1)
	nextMove := fmt.Sprintf("\"next-move\":\"%s-%s\"", fromSquare, toSquare)
	return nextMove
}

func createChessNode(board_json string) ChessNode {
	node := ChessNode{board: createBoard(board_json)}
	node.black_legal_moves = getLegalBlackMoves(node)
	return node
}

func createChessNodeUsingMap(dat map[string]string) ChessNode {
	node := ChessNode{board: createBoardUsingMap(dat)}
	node.black_legal_moves = getLegalBlackMoves(node)
	return node
}

func createChessNodeUsingArray(dat []string) ChessNode {
	node := ChessNode{board: createBoardUsingArray(dat)}
	node.black_legal_moves = getLegalBlackMoves(node)
	return node
}

func randomColumn() string {
	var columns string
	columns = "abcdefgh"
	return string(columns[rand.Intn(len(columns))])
}

//prototype #1 - choose a random initial pawn and move 2 places
func GetNextMove() string {
	var move string

	var randChar string
	randChar = randomColumn()
	move = fmt.Sprintf("\"next-move\":\"%s7-%s5\"", randChar, randChar)
	return move
}

//prototype #2 - calculate legal black moves and randomly choose a move
func GetNextMoveUsingArray(dat []string) string {
	var move string

	node := createChessNodeUsingArray(dat)
	//printNode(node)
	if len(node.black_legal_moves) > 0 {
		randMove := node.black_legal_moves[rand.Intn(len(node.black_legal_moves))]
		move = formatNextMove(randMove)
	}

	return move
}

func createBoardUsingArray(dat []string) [8][8]string {
	boardStr := dat[0] //should be first element in array
	boardLen := len(boardStr)

	//split on the '],'
	arr := strings.Split(boardStr[1:boardLen-1], "],")
	var board [8][8]string
	for i := 0; i < len(arr); i++ {
		row := arr[i][1:len(arr[i])] //remove the leading '['
		if i == 7 {                  //remove the trailing ']' on last row
			row = row[0 : len(row)-1]
		}
		//now split on the ','
		rowArr := strings.Split(row, ",")

		//finally populate matrix
		for j := 0; j < len(rowArr); j++ {
			if rowArr[j] == "0" {
				board[i][j] = rowArr[j]
			} else {
				//piece exists...strip off double quotes "xx"
				board[i][j] = rowArr[j][1:3]
			}
		}
	}
	return board
}

func createBoardUsingMap(dat map[string]string) [8][8]string {
	//populate board matrix
	var board [8][8]string
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			letter := numberToLetter(col)
			square := letter + strconv.Itoa(row+1)
			if val, ok := dat[square]; ok {
				board[row][col] = val
			} else {
				board[row][col] = "0"
			}
		}
	}
	return board
}

func convertJsonToMap(board_json string) map[string]string {
	//convert json to string map
	byt := []byte(board_json)
	var dat map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	return dat
}

//return an 8x8 board from the JSON representation
func createBoard(board_json string) [8][8]string {
	dat := convertJsonToMap(board_json)
	return createBoardUsingMap(dat)
}

func numberToLetter(x int) string {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	return letters[x]
}

func printNode(node ChessNode) {
	board := node.board
	for row := 7; row >= 0; row-- {
		output := ""
		for col := 0; col < 8; col++ {
			if board[row][col] == "0" {
				output += "\u3000"
			} else {
				output += pieceToUnicode(board[row][col])
			}
		}
		fmt.Println(output)
	}
}

func pieceToUnicode(piece string) string {
	return pieces[piece]["codepoint"]
}

//for the given ChessNode return all of the legal black moves
func getLegalBlackMoves(node ChessNode) []Move {
	var black_moves []Move
	board := node.board
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			//fmt.Println("piece=", piece)
			if piece == "0" {
				continue
			} else if piece[0:1] == "b" {
				//found a black piece
				//fmt.Println("black piece = ", piece, ", row=", row, "col=", col)
				black_moves = append(black_moves, getMovesForBlackPiece(piece, row, col, node)...)
			}
		}
	}
	//fmt.Println("# of legal black_moves = ", len(black_moves))
	return black_moves
}

func getMovesForBlackPiece(piece string, row int, col int, node ChessNode) []Move {
	var moves []Move
	piece_type := piece[1:2]
	switch piece_type {
	case "p":
		//fmt.Println("PAWN")
		moves = append(moves, getMovesForBlackPawn(piece, row, col, node)...)
	default:
		//fmt.Println("DEFAULT = ", piece)
	}
	return moves
}

func getMovesForBlackPawn(piece string, row int, col int, node ChessNode) []Move {
	var moves []Move

	//check if first square in front is blank
	if row > 0 && node.board[row-1][col] == "0" {
		moves = append(moves, Move{orig_row: row, orig_col: col, dest_row: row - 1, dest_col: col, piece: piece})

		//check 2 moves in front if on the initial row (6th for black)
		if row == 6 && node.board[row-2][col] == "0" {
			//check 2 moves in front
			moves = append(moves, Move{orig_row: row, orig_col: col, dest_row: row - 2, dest_col: col, piece: piece})
		}
	}

	//can you attack diagonally to the right?
	if col > 0 && row > 0 {
		attackSquare := node.board[row-1][col-1]
		if attackSquare != "0" && attackSquare[0:1] == "w" {
			moves = append(moves, Move{orig_row: row, orig_col: col, dest_row: row - 1, dest_col: col - 1, piece: piece})
		}
	}
	//can you attack diagonally to the left?
	if col < 7 && row > 0 {
		attackSquare := node.board[row-1][col+1]
		if attackSquare != "0" && attackSquare[0:1] == "w" {
			moves = append(moves, Move{orig_row: row, orig_col: col, dest_row: row - 1, dest_col: col + 1, piece: piece})
		}
	}
	return moves
}
