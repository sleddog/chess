package chess

type ChessNode struct {
	board        [8][8]string
	active_color string
	legal_moves  []Move
	depth        int
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
