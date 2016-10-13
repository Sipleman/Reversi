package GameField

import (
	"testing"
)

type testPair struct{
	function bool
	result bool
}

var field1 = [8][8]int{
{0, 0, 0, 0, 0, 0, 0, 0},
{0, 0, 1, 2, 2, 2, 1, 0},
{0, 0, 0, 2, 1, 2, 1, 0},
{0, 0, 0, 1, 0, 0, 0, 0},
{0, 0, 1, 2, 1, 0, 0, 0},
{0, 0, 0, 0, 0, 0, 0, 0},
{0, 0, 0, 0, 0, 0, 0, 0},
{0, 0, 0, 0, 0, 0, 0, 0},
}

var tests = []testPair{
	{getIsInFieldResult(1, 1), true},
	{getIsInFieldResult(0, 0), true},
	{getIsInFieldResult(-1, -1), false},
	{getIsInFieldResult(11, 11), false},
	{getIsInFieldResult(1, 15), false},
	{getIsInFieldResult(1123, -1), false},
	{getRecalculationResult(field1), true},
}

func TestGameField(t *testing.T){
	for _, pair := range tests {
		result := pair.function
		if result != pair.result{
			t.Error(
				"expected", pair.result,
				"got", result,
			)
		}
	}
}

func getIsInFieldResult(x, y int) bool{
	gameField := New()
	return gameField.isInField(x, y)
}
func getRecalculationResult(field [8][8]int) bool{
	gameField := New()
	gameField.Field = field
	gameField.recalculation()
	return gameField.player2_scores == calcScores(field, 2) && gameField.player1_scores == calcScores(field, 1)
}
func calcScores(field [8][8]int, player int) int{
	result := 0
	for i:=0; i<8; i++{
		for j:=0; j<8;j++{
			if field[i][j] == player {
				result++
			}
		}
	}
	return result
}
//func getMakeStepResult(){}