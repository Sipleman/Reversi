package GameField

import tl "github.com/JoelOtter/termloop"

type Player struct {
	Entity *tl.Entity
	rectangle *tl.Rectangle
	playerNumber int
}
