package controller

import "github.com/hajimehoshi/ebiten"

type Controller struct {

}

func New() *Controller {
	c := Controller{}

	return &c
}

func (c Controller) IsLeftPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
}

func (c Controller) IsRightPressed() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)
}