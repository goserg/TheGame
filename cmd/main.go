package main

import (
	"github.com/goserg/plat/game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	g, err := game.New()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
