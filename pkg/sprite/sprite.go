package sprite

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/odedro987/mig-engine/pkg/graphic"
	"github.com/odedro987/mig-engine/pkg/math"
	"github.com/odedro987/mig-engine/pkg/object"
)

type Base struct {
	object.Base
	graphic *graphic.GxlGraphic
}

func New(x, y float64) GxlSprite {
	s := &Base{}
	s.SetPosition(x, y)
	return s
}

func (s *Base) MakeGraphic(w, h int, c color.Color) {
	s.graphic = graphic.MakeGraphic(w, h, c)
	s.SetSize(w, h)
}

func (s *Base) LoadGraphic(path string) {
	s.graphic = graphic.LoadGraphic(path)
	w, h := s.graphic.GetImage().Size()
	s.SetSize(w, h)
}

func (s *Base) Draw(screen *ebiten.Image) {
	if s.graphic == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(s.Angle)
	op.GeoM.Translate(s.X, s.Y)
	if *s.Scale != *math.NewPoint(1, 1) {
		op.GeoM.Scale(s.Scale.X, s.Scale.Y)
	}
	screen.DrawImage(s.graphic.GetImage(), op)
}

type GxlSprite interface {
	object.GxlObject
	MakeGraphic(w, h int, c color.Color)
}