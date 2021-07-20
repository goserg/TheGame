package player

import (
	"github.com/goserg/plat/animator"
	"github.com/goserg/plat/game/controller"
	"github.com/goserg/plat/game/physics"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"os"
	"strconv"
)

type Player struct {
	animator.Animator
	controller.Controller

	img *ebiten.Image
	idleAnim []*ebiten.Image
	idleAnimState int

	position *physics.Vector
	speed *physics.Vector
	maxSpeed *physics.Vector
}

func (p *Player) Update() {
	p.img, _ = p.GetFrame(animator.Idle)

	p.speed.X = 0
	if p.IsLeftPressed() {
		p.speed.X -= p.maxSpeed.X
	}
	if p.IsRightPressed() {
		p.speed.X += p.maxSpeed.X
	}

	p.position.Add(p.speed)
}

func (p *Player) Draw(screen *ebiten.Image) {
	var geoM ebiten.GeoM
	geoM.Translate(p.position.X, p.position.Y)
	err := screen.DrawImage(p.img, &ebiten.DrawImageOptions{GeoM: geoM})
	if err != nil {
		panic(err)
	}
}

func New() (*Player, error) {
	p := Player{
		Animator: animator.New(),
		position: &physics.Vector{},
		speed:         &physics.Vector{},
		maxSpeed:      &physics.Vector{
			X: 4,
			Y: 10,
		},
	}

	for i := 1; i < 10; i++ {
		f, err := os.Open("assets/player/Gino Character/PNG/Idle, run, jump/idle0" + strconv.Itoa(i) + ".png")
		if err != nil {
			return nil, err
		}
		img, _, err := image.Decode(f)
		if err != nil {
			return nil, err
		}
		ei, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		p.Animator.AddFrame(animator.Idle, ei)
	}
	err := p.Play(animator.Idle)
	if err != nil {
		return nil, err
	}
	return &p, nil
}