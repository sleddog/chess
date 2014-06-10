package chess

func getMoves(color string, board [8][8]string) []Move {
	var moves []Move
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			piece := board[row][col]
			if piece == "0" {
				continue
			} else if piece[0:1] == color {
				moves = append(moves, getPieceMoves(color, piece, row, col, board)...)
			}
		}
	}
	return moves
}

func getPieceMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var moves []Move
	piece_type := piece[1:2]
	switch piece_type {
	case "p":
		moves = append(moves, getPawnMoves(color, piece, row, col, board)...)
	case "n":
		moves = append(moves, getKnightMoves(color, piece, row, col, board)...)
	case "b":
		moves = append(moves, getBishopMoves(color, piece, row, col, board)...)
	case "r":
		moves = append(moves, getRookMoves(color, piece, row, col, board)...)
	case "q":
		moves = append(moves, getQueenMoves(color, piece, row, col, board)...)
	case "k":
		moves = append(moves, getKingMoves(color, piece, row, col, board)...)
	default:
	}
	return moves
}

func getPawnMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var moves []Move
	fromCoord := Coord{row: row, col: col}

	if color == "b" {
		//check if first square in front is blank
		if row > 0 && board[row-1][col] == "0" {
			moves = append(moves,
				Move{from: fromCoord,
					to:    Coord{row: row - 1, col: col},
					piece: piece})

			//check 2 moves in front if on the initial row (6th for black)
			if row == 6 && board[row-2][col] == "0" {
				//check 2 moves in front
				moves = append(moves,
					Move{from: fromCoord,
						to:    Coord{row: row - 2, col: col},
						piece: piece})
			}
		}

		//can you attack diagonally to the right?
		if col > 0 && row > 0 {
			attackSquare := board[row-1][col-1]
			if attackSquare != "0" && attackSquare[0:1] == opposite(color) {
				moves = append(moves,
					Move{from: fromCoord,
						to:    Coord{row: row - 1, col: col - 1},
						piece: piece})
			}
		}
		//can you attack diagonally to the left?
		if col < 7 && row > 0 {
			attackSquare := board[row-1][col+1]
			if attackSquare != "0" && attackSquare[0:1] == opposite(color) {
				moves = append(moves,
					Move{from: fromCoord,
						to:    Coord{row: row - 1, col: col + 1},
						piece: piece})
			}
		}

		//TODO en passant

	} else { //color == "w"
		//TODO white pawn moves...
	}
	return moves
}

func canMove(color string, board [8][8]string, move Coord) bool {
	//is off board?
	if move.col < 0 || move.col > 7 || move.row < 0 || move.row > 7 {
		return false
	}
	square := board[move.row][move.col]
	//is the square empty?
	if square == "0" {
		return true
	} else { //square is occupied
		if square[0:1] == opposite(color) {
			//opposite piece exists in this square
			return true
		} else {
			//friendly piece is already in place here
			return false
		}
	}
}

func getKnightMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var moves []Move
	var possibleMoves [8]Coord
	possibleMoves[0] = Coord{row: row + 2, col: col - 1}
	possibleMoves[1] = Coord{row: row + 2, col: col + 1}
	possibleMoves[2] = Coord{row: row - 2, col: col - 1}
	possibleMoves[3] = Coord{row: row - 2, col: col + 1}
	possibleMoves[4] = Coord{row: row + 1, col: col - 2}
	possibleMoves[5] = Coord{row: row + 1, col: col + 2}
	possibleMoves[6] = Coord{row: row - 1, col: col - 2}
	possibleMoves[7] = Coord{row: row - 1, col: col + 2}

	fromCoord := Coord{row: row, col: col}
	for i := 0; i < 8; i++ {
		if canMove(color, board, possibleMoves[i]) {
			moves = append(moves,
				Move{from: fromCoord,
					to:    possibleMoves[i],
					piece: piece})
		}
	}
	return moves
}

func getDirectionalMoves(color string, board [8][8]string, dirs [][2]int, from Coord, piece string) []Move {
	var moves []Move
	for i := 0; i < len(dirs); i++ {
		dir := dirs[i]
		step := 1
		//for each direction, travel until the edge of board or piece
		for {
			move := Coord{row: from.row + dir[0]*step,
				col: from.col + dir[1]*step}
			if canMove(color, board, move) {
				moves = append(moves, Move{from: from, to: move, piece: piece})
				step = step + 1
				//if move is a opposite piece, stop calculating on this line
				square := board[move.row][move.col]
				if square != "0" && square[0:1] == opposite(color) {
					break
				}
			} else {
				break
			}
		}
	}
	return moves
}

func getBishopMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var directions [][2]int
	directions = append(directions, [2]int{1, 1})
	directions = append(directions, [2]int{1, -1})
	directions = append(directions, [2]int{-1, 1})
	directions = append(directions, [2]int{-1, -1})
	from := Coord{row: row, col: col}
	return getDirectionalMoves(color, board, directions, from, piece)
}

func getRookMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var directions [][2]int
	directions = append(directions, [2]int{1, 0})
	directions = append(directions, [2]int{0, 1})
	directions = append(directions, [2]int{-1, 0})
	directions = append(directions, [2]int{0, -1})
	from := Coord{row: row, col: col}
	return getDirectionalMoves(color, board, directions, from, piece)
}

func getQueenMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var moves []Move
	moves = append(moves, getBishopMoves(color, piece, row, col, board)...)
	moves = append(moves, getRookMoves(color, piece, row, col, board)...)
	return moves
}

func getKingMoves(color string, piece string, row int, col int, board [8][8]string) []Move {
	var moves []Move
	var possibleMoves [8]Coord
	possibleMoves[0] = Coord{row: row + 1, col: col}
	possibleMoves[1] = Coord{row: row, col: col + 1}
	possibleMoves[2] = Coord{row: row - 1, col: col}
	possibleMoves[3] = Coord{row: row, col: col - 1}
	possibleMoves[4] = Coord{row: row + 1, col: col + 1}
	possibleMoves[5] = Coord{row: row + 1, col: col - 1}
	possibleMoves[6] = Coord{row: row - 1, col: col - 1}
	possibleMoves[7] = Coord{row: row - 1, col: col + 1}

	fromCoord := Coord{row: row, col: col}
	for i := 0; i < 8; i++ {
		if canMove(color, board, possibleMoves[i]) {
			moves = append(moves,
				Move{from: fromCoord,
					to:    possibleMoves[i],
					piece: piece})
		}
	}
	return moves
}

func opposite(color string) string {
	if color == "b" {
		return "w"
	} else {
		return "b"
	}
}
