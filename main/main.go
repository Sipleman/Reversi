package main

import tl "github.com/JoelOtter/termloop"
import "./GameField"

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Fg: tl.ColorBlack,
		Ch: ' ',
	})
	//filed := GameField{level: tl.NewBaseLevel(tl.Cell{
	//	Bg: tl.ColorMagenta,
	//	Fg: tl.ColorYellow,
	//	Ch: '-',
	//})}
	//filed.level.AddEntity(tl.NewRectangle(10,10,30,30,tl.ColorBlack))
	//

	//// Set the character at position (0, 0) on the entity.
	//level.AddEntity(&player)
	//level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))
	//level.AddEntity(tl.NewText(2,2,"looool", tl.ColorWhite, 0))
	gameField := GameField.New()
	level.AddEntity(&gameField)
	game.Screen().SetLevel(level)
	game.Start()
}
