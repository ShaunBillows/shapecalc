package tests

import (
	"github.com/ShaunBillows/shapes-cli-project-go/app/shapes"
	"github.com/ShaunBillows/shapes-cli-project-go/internal/messages"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        shapes.Shape
		hasPerimeter float64
		err          error
	}{
		{name: "circle test: positive radius", shape: shapes.Circle{Radius: 10}, hasPerimeter: 62.83185307179586, err: nil},
		{name: "circle test: negative radius", shape: shapes.Circle{Radius: -10}, hasPerimeter: 0, err: messages.ErrNegativeRadius},
		{name: "rectangle test: positive width and height", shape: shapes.Rectangle{Height: 10, Width: 10}, hasPerimeter: 40, err: nil},
		{name: "rectangle test: negative width", shape: shapes.Rectangle{Height: 10, Width: -10}, hasPerimeter: 0, err: messages.ErrNegativeWidth},
		{name: "rectangle test: negative height", shape: shapes.Rectangle{Height: -10, Width: 10}, hasPerimeter: 0, err: messages.ErrNegativeHeight},
		{name: "rectangle test: negative width and height", shape: shapes.Rectangle{Height: -10, Width: -10}, hasPerimeter: 0, err: messages.ErrNegativeHeight},
		{name: "triangle test: positive base and height", shape: shapes.Triangle{Height: 3, Base: 5}, hasPerimeter: 12.810249675906654, err: nil},
		{name: "triangle test: negative base", shape: shapes.Triangle{Height: 10, Base: -10}, hasPerimeter: 0, err: messages.ErrNegativeBase},
		{name: "triangle test: negative base and height", shape: shapes.Triangle{Height: -10, Base: -10}, hasPerimeter: 0, err: messages.ErrNegativeHeight},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shape.Perimeter()
			assertError(t, tt.name, err, tt.err)
			assertFloat64(t, tt.name, got, tt.hasPerimeter)
		})
	}
}
