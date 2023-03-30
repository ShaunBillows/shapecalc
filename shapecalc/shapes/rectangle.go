package shapes

import (
	"errors"
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
		return 0, errors.New(ErrNegativeHeight)
	}
	if r.Width < 0 {
		return 0, errors.New(ErrNegativeWidth)
	}
	return r.Height * r.Width, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.Height < 0 {
		return 0, errors.New(ErrNegativeHeight)
	}
	if r.Width < 0 {
		return 0, errors.New(ErrNegativeWidth)
	}
	return 2*r.Height + 2*r.Width, nil
}

func (r Rectangle) Type() ShapeType {
	return ShapeTypeRectangle
}
