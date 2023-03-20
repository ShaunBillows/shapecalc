package shapes

import (
	"github.com/ShaunBillows/shapes-cli-project-go/internal/messages"
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
		return 0, messages.ErrNegativeRadius
	}
	return math.Pi * c.Radius * c.Radius, nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.Radius < 0 {
		return 0, messages.ErrNegativeRadius
	}
	return 2 * math.Pi * c.Radius, nil
}
