package entities

import (
	"image/color"

	"github.com/odedro987/gixel-engine/examples/testing/systems"
	"github.com/odedro987/gixel-engine/gixel"
	"github.com/odedro987/gixel-engine/gixel/graphic"
	"github.com/odedro987/gixel-engine/gixel/systems/flipping"
	"github.com/odedro987/gixel-engine/gixel/systems/physics"
)

type Player struct {
	gixel.BaseGxlSprite
	// Systems
	flipping.Flipping
	physics.Physics
	systems.Movement
}

func NewPlayer(x, y float64) *Player {
	p := &Player{}
	p.SetPosition(x, y)

	return p
}

func (p *Player) Init(game *gixel.GxlGame) {
	p.BaseGxlSprite.Init(game)

	p.Flipping.Init(p)
	p.Physics.Init(p)
	p.Movement.Init(p)

	p.ApplyGraphic(game.Graphics().MakeGraphic(16, 16, color.White, graphic.CacheOptions{}))
}

func (p *Player) Update(elapsed float64) error {
	err := p.BaseGxlSprite.Update(elapsed)
	if err != nil {
		return err
	}

	p.Flipping.Update()
	p.Physics.Update(elapsed)
	p.Movement.Update(elapsed)

	return nil
}
