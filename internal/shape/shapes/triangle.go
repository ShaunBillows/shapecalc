package shapes

import (
	"github.com/ShaunBillows/shapes-cli-project-go/internal/messages"
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
		return 0, messages.ErrNegativeHeight
	}
	if t.Base < 0 {
		return 0, messages.ErrNegativeBase
	}
	return 0.5 * t.Base * t.Height, nil
}

func (t Triangle) Perimeter() (float64, error) {
	if t.Height < 0 {
		return 0, messages.ErrNegativeHeight
	}
	if t.Base < 0 {
		return 0, messages.ErrNegativeBase
	}
	return 2*math.Sqrt(t.Height*t.Height+math.Pow(t.Base/2, 2)) + t.Base, nil
}
