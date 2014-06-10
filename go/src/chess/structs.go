package chess

type ChessNode struct {
	board             [8][8]string
	white_legal_moves []Move
	black_legal_moves []Move
}

type Move struct {
	from  Coord
	to    Coord
	piece string
	value int
}

type Coord struct {
	row int
	col int
}
