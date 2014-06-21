package chess

import (
	//"fmt"
	"math/rand"
)

//Minimax-Decision(state) returns an action
//v = Max-Value(state)
//return the action in successors(state) with value v
func miniMaxDecision(state ChessNode) string {
	//fmt.Println("miniMaxDecision(state=", state, ")")
	//fmt.Println("pre utility = ", state.utility_value)
	v := minValue(state)
	//fmt.Println("v=", v)
	moves := successors(state)
	var equalMoves []Move
	for i := 0; i < len(moves); i++ {
		//ns := nextState(state, moves[i])
		//ns.active_color = state.active_color
		//fmt.Println("moves[", i, "]", moves[i])
		newBoard := makeMove(state.board, moves[i])
		u := utility(opposite(state.active_color), newBoard)
		//fmt.Println("u=", u)
		if v == u {
			//fmt.Println("FOUND, move =", moves[i])
			equalMoves = append(equalMoves, moves[i])
		}
	}
	move := ""
	if equalMoves != nil {
		randMove := equalMoves[rand.Intn(len(equalMoves))]
		move = formatNextMove(randMove)
	} else {
		//somehow couldn't find a move that matched max, so just grab one
		randMove := moves[rand.Intn(len(moves))]
		move = formatNextMove(randMove)
	}
	return move
}

//Max-Value(state) returns a utility value
//If Terminal-Test(state) then return Utility(state)
//v <= -infinity
//for a, s in Successors(state) do
//  v <= Max(v, Min-Value(s))
//return v
func maxValue(state ChessNode) int {
	if terminalTest(state) {
		return state.utility_value
	}

	v := -9999999

	moves := successors(state)
	for i := 0; i < len(moves); i++ {
		s := minValue(nextState(state, moves[i]))
		if s >= v {
			//fmt.Println("\n\n----->found larger value, ", s, ",", moves[i], "\n\n")
			v = s
		}
	}
	return v
}

//Min-Value(state) returns a utility value
//If Terminal-Test(state) then return Utility(state)
//v <= +infinity
//for a, s in Successors(state) do
//  v <= Min(v, Max-Value(s))
//return v
func minValue(state ChessNode) int {
	//fmt.Println("minValue, state=", state)
	if terminalTest(state) {
		return state.utility_value
	}

	v := 9999999
	moves := successors(state)
	//state.active_color = opposite(state.active_color)
	for i := 0; i < len(moves); i++ {
		s := maxValue(nextState(state, moves[i]))
		if s <= v {
			//fmt.Println("---->found smaller value, ", s, ",", moves[i])
			v = s
		}
	}
	return v
}

func terminalTest(state ChessNode) bool {
	//TODO check for checkmate
	//stop at a certain depth, then return the utility
	if state.depth >= 1 {
		return true
	}
	return false
}

func successors(state ChessNode) []Move {
	var moves []Move
	moves = getLegalMoves(state.active_color, state.board)
	return moves
}

func utility(color string, board [8][8]string) int {
	return calculatePointValue(color, board) - calculatePointValue(opposite(color), board)
}

func nextState(state ChessNode, move Move) ChessNode {
	newBoard := makeMove(state.board, move)
	color := opposite(state.active_color)
	return ChessNode{
		board:         newBoard,
		active_color:  color,
		depth:         state.depth + 1,
		prev_move:     move,
		utility_value: utility(color, newBoard),
	}
}
