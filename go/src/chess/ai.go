package chess

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

//Get Next move using mini max algorithm
func GetNextMoveUsingMiniMax(dat []string, prev_move string) string {
	var move string
	var stats string
	active_color := "b"
	node := createChessNodeUsingArray(dat, active_color, prev_move)
	move, stats = miniMaxDecision(node)
	return move + "," + stats
}

// Get Next Move based on minimizing opponent material value
// - calculate all possible legal moves
// - construct the resulting board position for each move
// - sum the material value of each piece
// - return the move that minimizes the total material value of opponent
func GetNextMoveUsingPointValue(dat []string) string {
	var move string

	node := createChessNodeUsingArray(dat, "", "")
	node.legal_moves = getLegalMoves("b", node.board)
	//printNode(node)

	if len(node.legal_moves) > 0 {
		for i := 0; i < len(node.legal_moves); i++ {
			newBoard := makeMove(node.board, node.legal_moves[i])
			materialValue := calculatePointValue("w", newBoard)
			node.legal_moves[i].value = materialValue
		}

		moves := node.legal_moves
		sort.Sort(ByMaterialValue(moves))

		//after sorting, choose a random move that has the same minimum score
		minMove := moves[0]
		var j int
		for j = 0; j < len(moves); j++ {
			if moves[j].value > minMove.value {
				break
			}
		}
		randMove := node.legal_moves[rand.Intn(j)]
		move = formatNextMove(randMove)
	}

	return move
}

//define global variables
var pieces = make(map[string]map[string]string)
var initial_board_json string
var number_of_nodes_per_depth = make(map[string]int)

func init() {
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

	//set the initial_board_json
	initial_board_json = "{\"a8\":\"br\",\"b8\":\"bn\",\"c8\":\"bb\",\"d8\":\"bq\",\"e8\":\"bk\",\"f8\":\"bb\",\"g8\":\"bn\",\"h8\":\"br\",\"a7\":\"bp\",\"b7\":\"bp\",\"c7\":\"bp\",\"d7\":\"bp\",\"e7\":\"bp\",\"f7\":\"bp\",\"g7\":\"bp\",\"h7\":\"bp\",\"a2\":\"wp\",\"b2\":\"wp\",\"c2\":\"wp\",\"d2\":\"wp\",\"e2\":\"wp\",\"f2\":\"wp\",\"g2\":\"wp\",\"h2\":\"wp\",\"a1\":\"wr\",\"b1\":\"wn\",\"c1\":\"wb\",\"d1\":\"wq\",\"e1\":\"wk\",\"f1\":\"wb\",\"g1\":\"wn\",\"h1\":\"wr\"}"
}

func pieceToUnicode(piece string) string {
	return pieces[piece]["codepoint"]
}

func formatNextMove(move Move) string {
	columns := "abcdefgh"
	fromSquare := string(columns[move.from.col]) + strconv.Itoa(move.from.row+1)
	toSquare := string(columns[move.to.col]) + strconv.Itoa(move.to.row+1)
	nextMove := fmt.Sprintf("\"next-move\":\"%s-%s\"", fromSquare, toSquare)
	return nextMove
}

func updateStats(state ChessNode, count int) {
	key := strconv.Itoa(state.depth) + opposite(state.active_color)
	number_of_nodes_per_depth[key] += count
}

func formatStats() string {
	b, err := json.Marshal(number_of_nodes_per_depth)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	stats := fmt.Sprintf("\"stats\":%s", string(b))
	return stats
}

func calculatePointValue(color string, board [8][8]string) int {
	//sum up the point values of all of the chess pieces for this color
	sum := 0
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == "0" {
				continue
			} else if piece[0:1] == color {
				sum = sum + pieceValue(piece[1:2])
			}
		}
	}
	return sum
}

func pieceValue(piece_type string) int {
	switch piece_type {
	case "p":
		return 1
	case "n":
		return 3
	case "b":
		return 3
	case "r":
		return 5
	case "q":
		return 9
	default:
		return 0
	}
}

type ByMaterialValue []Move

func (a ByMaterialValue) Len() int           { return len(a) }
func (a ByMaterialValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMaterialValue) Less(i, j int) bool { return a[i].value < a[j].value }
