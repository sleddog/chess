package chess

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func createState(board [8][8]string, active_color string, prev_move string) ChessNode {
	node := ChessNode{
		board:         board,
		depth:         0,
		active_color:  active_color,
		prev_move:     stringToMove(prev_move, board),
		utility_value: utility(active_color, board),
	}
	return node
}

//"e2-e4" to Move{}
func stringToMove(str string, board [8][8]string) Move {
	move := Move{
		from: stringToCoord(str[0:2]),
		to:   stringToCoord(str[3:5]),
	}
	//determine piece by checking the board at the 'to' position
	move.piece = board[move.to.row][move.to.col]
	return move
}

func stringToCoord(str string) Coord {
	letters := "abcdefgh"
	row := strings.Index(letters, str[0:1])
	if intval, err := strconv.Atoi(str[1:2]); err == nil {
		col := intval - 1
		return Coord{row: row, col: col}
	}
	return Coord{row: -1, col: -1}
}

func createChessNodeUsingArray(dat []string, active_color, prev_move string) ChessNode {
	board := createBoardUsingArray(dat)
	return createState(board, active_color, prev_move)
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

func createChessNodeUsingMap(dat map[string]string) ChessNode {
	node := ChessNode{board: createBoardUsingMap(dat)}
	node.legal_moves = getMoves("b", node.board)
	return node
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

func createBoardUsingFen(fen string) [8][8]string {
	//example = 'rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR'
	rows := strings.Split(fen, "/")
	currentRow := 8
	var boardMap map[string]string
	boardMap = make(map[string]string)
	var step int
	for i := 0; i < len(rows); i++ {
		currentCol := 1
		row := rows[i]
		for j := 0; j < len(row); j++ {
			char := string(row[j])
			if intval, err := strconv.Atoi(char); err == nil {
				step = intval
			} else { //add the piece
				l := numberToLetter(currentCol - 1)
				square := l + strconv.Itoa(currentRow)
				piece := convertFenPiece(char)
				boardMap[square] = piece
				step = 1
			}
			currentCol = currentCol + step
		}
		currentRow = currentRow - 1
	}
	//now convert the map into the board
	return createBoardUsingMap(boardMap)
}

//converts FEN piece to my representation, i.e. 'K'->'wk', 'q'->'bq'
func convertFenPiece(char string) string {
	//uppercase character means white, otherwise black
	var piece string
	if strings.ToUpper(char) == char {
		piece = "w"
	} else {
		piece = "b"
	}
	return piece + strings.ToLower(char)
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

func createChessNode(board_json string) ChessNode {
	node := ChessNode{board: createBoard(board_json)}
	node.legal_moves = getMoves("b", node.board)
	return node
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
