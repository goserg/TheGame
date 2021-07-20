package physics

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
)

type Collider struct {
	Position Vector
	Size     Vector
}

func (c *Collider) Draw(screen *ebiten.Image) error {
	img, err := ebiten.NewImage(int(c.Size.X), int(c.Size.Y), ebiten.FilterDefault)
	if err != nil {
		return err
	}
	if err := img.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 100}); err != nil {
		return err
	}

	var geoM ebiten.GeoM
	geoM.Translate(c.Position.X, c.Position.Y)
	if err := screen.DrawImage(img, &ebiten.DrawImageOptions{GeoM: geoM}); err != nil {
		return err
	}
	return nil
}
