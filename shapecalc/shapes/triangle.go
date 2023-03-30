package shapes

import (
	"errors"
	"math"
)

type Triangle struct {
	Base   float64
	Height float64
}

func NewTriangle() *Triangle {
	return &Triangle{}
}

func (t Triangle) Area() (float64, error) {
	if t.Height < 0 {
		return 0, errors.New(ErrNegativeHeight)
	}
	if t.Base < 0 {
		return 0, errors.New(ErrNegativeBase)
	}
	return 0.5 * t.Base * t.Height, nil
}

func (t Triangle) Perimeter() (float64, error) {
	if t.Height < 0 {
		return 0, errors.New(ErrNegativeHeight)
	}
	if t.Base < 0 {
		return 0, errors.New(ErrNegativeBase)
	}
	return 2*math.Sqrt(t.Height*t.Height+math.Pow(t.Base/2, 2)) + t.Base, nil
}

func (r Triangle) Type() ShapeType {
	return ShapeTypeTriangle
}
