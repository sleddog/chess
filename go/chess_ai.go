package chess

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
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
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	//populate board matrix
	var board [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			letter := numberToLetter(j)
			square := letter + strconv.Itoa(i+1)
			if val, ok := dat[square].(string); ok {
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
