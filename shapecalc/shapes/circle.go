package shapes

import (
	"errors"
	"math"
)

type Circle struct {
	Radius float64
}

func NewCircle() *Circle {
	return &Circle{}
}

func (c Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New(ErrNegativeRadius)
	}
	return math.Pi * c.Radius * c.Radius, nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New(ErrNegativeRadius)
	}
	return 2 * math.Pi * c.Radius, nil
}

func (r Circle) Type() ShapeType {
	return ShapeTypeCircle
}
