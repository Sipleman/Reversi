package main

import tl "github.com/JoelOtter/termloop"
import "./GameField"

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	gameField := GameField.New()
	level.AddEntity(&gameField)
	game.Screen().SetLevel(level)

	game.Start()
}
