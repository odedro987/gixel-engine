package gixel

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GxlGroup[T GxlBasic] struct {
	BaseGxlBasic
	members []*T
	length  int
	maxSize int
}

func NewGroup[T GxlBasic](maxSize int) *GxlGroup[T] {
	if maxSize < 0 {
		panic("maxSize cannot be negative")
	}

	var members []*T
	if maxSize == 0 {
		members = make([]*T, 0)
	} else {
		members = make([]*T, maxSize)
	}

	group := GxlGroup[T]{
		members: members,
		maxSize: maxSize,
	}

	return &group
}

func (g *GxlGroup[T]) Add(object T) *T {
	freeSlotIdx := -1
	for idx, member := range g.members {
		if member == &object {
			//warn
			return nil
		}

		if freeSlotIdx == -1 && member == nil {
			freeSlotIdx = idx
		}
	}

	if freeSlotIdx != -1 {
		g.members[freeSlotIdx] = &object

		if freeSlotIdx >= g.length {
			g.length = freeSlotIdx + 1
		}

		(object).Init()
		return &object
	}

	if g.maxSize > 0 && g.length >= g.maxSize {
		//warn
		return nil
	}

	g.members = append(g.members, &object)
	g.length++

	object.Init()
	return &object
}

func (g *GxlGroup[T]) Remove(object *T) *T {
	for idx, member := range g.members {
		if member == object {
			g.members[idx] = nil
			return object
		}
	}

	return nil
}

func (g *GxlGroup[T]) Range(f func(idx int, value *T) bool) {
	for idx, m := range g.members {
		if m == nil {
			continue
		}

		if !f(idx, m) {
			break
		}
	}
}

func (g *GxlGroup[T]) Draw(screen *ebiten.Image) {
	for _, m := range g.members {
		if m != nil && (*m).Exists() && (*m).IsVisible() {
			(*m).Draw(screen)
		}
	}
}

func (g *GxlGroup[T]) Update(elapsed float64) error {
	for _, m := range g.members {
		if m != nil && (*m).Exists() {
			err := (*m).Update(elapsed)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *GxlGroup[T]) Destroy() {
	for _, m := range g.members {
		if m != nil {
			(*m).Destroy()
		}
	}
}
