package chess

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var pieces = make(map[string]map[string]string)
//pieces = {
//  "wk": {"html": "&#9812;","codepoint":"\u2654"},
//  "wq": {"html": "&#9813;","codepoint":"\u2655"},
//  "wr": {"html": "&#9814;","codepoint":"\u2656"},
//  "wb": {"html": "&#9815;","codepoint":"\u2657"},
//  "wn": {"html": "&#9816;","codepoint":"\u2658"},
//  "wp": {"html": "&#9817;","codepoint":"\u2659"},
//  "bk": {"html": "&#9818;","codepoint":"\u265A"},
//  "bq": {"html": "&#9819;","codepoint":"\u265B"},
//  "br": {"html": "&#9820;","codepoint":"\u265C"},
//  "bb": {"html": "&#9821;","codepoint":"\u265D"},
//  "bn": {"html": "&#9822;","codepoint":"\u265E"},
//  "bp": {"html": "&#9823;","codepoint":"\u265F"}}

func init() {
  fmt.Println("Inside init()")
	rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("pieces = ", pieces)
}

type Board struct {
  squares [8][8]string
  white_legal_moves []string
  black_legal_moves []string
}

func createBoard(board_json string) Board {
  b := Board{squares:initBoard(board_json)}
  fmt.Println("Default board is: ", b)
  return b
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

func initBoard(board_json string) [8][8]string {
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

func printBoard(board Board) {
	for i := 0; i < 8; i++ {
    var row string
		for j := 0; j < 8; j++ {
      if board.squares[i][j] == "0" {
          row = "\u3000"
      } else {
          row = "asdf" //pieceToUnicode(board[i][j])
      }
		}
    fmt.Println(row)
	}
}

func pieceToUnicode(piece string) string{
  return "blahhh"
}
