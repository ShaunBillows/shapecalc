package shapes

import (
	"github.com/ShaunBillows/shapes-cli-project-go/internal/messages"
)

type Rectangle struct {
	Height float64
	Width  float64
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}

func (r Rectangle) Area() (float64, error) {
	if r.Height < 0 {
		return 0, messages.ErrNegativeHeight
	}
	if r.Width < 0 {
		return 0, messages.ErrNegativeWidth
	}
	return r.Height * r.Width, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.Height < 0 {
		return 0, messages.ErrNegativeHeight
	}
	if r.Width < 0 {
		return 0, messages.ErrNegativeWidth
	}
	return 2*r.Height + 2*r.Width, nil
}

func (r Rectangle) Type() ShapeType {
	return ShapeTypeRectangle
}
