package chess

import (
	"fmt"
	"testing"
)

func TestGetNextMove(t *testing.T) {
   move := getNextMove()
   fmt.Println(move)
}
