package animator

import (
	"errors"
	"github.com/hajimehoshi/ebiten"
	"time"
)

type tag string

const (
	Idle tag = "idle"
)

type Animator map[tag]*Animation

func New() Animator {
	anim := make(map[tag]*Animation)
	return anim
}

type Animation struct {
	frames []*ebiten.Image
	n int
	delay time.Duration
	lastSwitch time.Time
}

func (m Animator) AddFrame(t tag, img *ebiten.Image) {
	_, ex := m[t]
	if !ex {
		m[t] = &Animation{
			delay: time.Second / 10,
		}
	}
	m[t].frames = append(m[t].frames, img)
}

func (m Animator) GetNextFrame(t tag) *ebiten.Image {
	frame := m[t].frames[m[t].n]
	m[t].n++
	if m[t].n >= len(m[t].frames) {
		m[t].n = 0
	}
	return frame
}

func (m Animator) GetFrame(t tag) (*ebiten.Image, error) {
	anim, ex := m[t]
	if !ex {
		return nil, errors.New(`no animation with tag "` + string(t) + `"`)
	}
	return anim.frames[anim.n], nil
}

func (m Animator) Play(t tag) error {
	anim, ex := m[t]
	if !ex {
		return errors.New(`no animation with tag "` + string(t) + `"`)
	}
	go func() {
		for {
			if time.Since(anim.lastSwitch) >= anim.delay {
				anim.lastSwitch = time.Now()
				anim.n++
				if anim.n >= len(anim.frames) {
					anim.n = 0
				}
			}
		}
	}()
	return nil
}
