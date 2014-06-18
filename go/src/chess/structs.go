package chess

type ChessNode struct {
	board         [8][8]string
	active_color  string
	prev_move     Move
	legal_moves   []Move
	depth         int
	utility_value int
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
