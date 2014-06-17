package chess

import (
	"math/rand"
)

//Minimax-Decision(state) returns an action
//v = Max-Value(state)
//return the action in successors(state) with value v
func miniMaxDecision(state ChessNode) string {
	v := maxValue(state)
	moves := successors(state)
	var equalMoves []Move
	for i := 0; i < len(moves); i++ {
		u := utility(nextState(state, moves[i]))
		if v == u {
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
		return utility(state)
	}

	v := -9999999

	moves := successors(state)
	for i := 0; i < len(moves); i++ {
		s := minValue(nextState(state, moves[i]))
		if s >= v {
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
	if terminalTest(state) {
		return utility(state)
	}

	v := 9999999
	moves := successors(state)
	for i := 0; i < len(moves); i++ {
		s := maxValue(nextState(state, moves[i]))
		if s <= v {
			v = s
		}
	}
	return v
}

func terminalTest(state ChessNode) bool {
	//TODO check for checkmate
	//stop at a certain depth, then return the utility
	if state.depth >= 2 {
		return true
	}
	return false
}

func successors(state ChessNode) []Move {
	var moves []Move
	moves = getLegalMoves(state.active_color, state.board)
	return moves
}

func utility(state ChessNode) int {
	utilityValue := calculatePointValue(state.active_color, state.board) -
		calculatePointValue(opposite(state.active_color), state.board)
	return utilityValue
}

func nextState(state ChessNode, move Move) ChessNode {
	return ChessNode{
		board:        makeMove(state.board, move),
		active_color: opposite(state.active_color),
		depth:        state.depth + 1,
	}
}
