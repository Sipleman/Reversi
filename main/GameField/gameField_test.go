package GameField

import "testing"

type testPair struct{
	function bool
	result bool
}

var tests = []testPair{
	{getIsInFieldResult(1, 1), true},
	{getIsInFieldResult(0, 0), true},
	{getIsInFieldResult(-1, -1), false},
	{getIsInFieldResult(11, 11), false},
	{getIsInFieldResult(1, 15), false},
	{getIsInFieldResult(1123, -1), false},
	{getRecalculationResult(), true},
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
func getRecalculationResult() bool{
	gameField := New()
	field := [8][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 2, 2, 2, 1, 0},
		{0, 0, 0, 2, 1, 2, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 1, 2, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	gameField.Field = field
	gameField.recalculation()
	return gameField.player2_scores == 6 && gameField.player1_scores == 7
}
//func getMakeStepResult(){}