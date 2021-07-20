package game

import (
	"github.com/goserg/plat/game/player"
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	player *player.Player
}

func (g *Game) Update(_ *ebiten.Image) error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return outsideWidth / 2, outsideHeight / 2
}

func New() (*Game, error) {
	var game Game
	p, err := player.New()
	if err != nil {
		return nil, err
	}
	game.player = p
	return &game, nil
}
