package tests

import (
	"github.com/ShaunBillows/shapes-cli-project-go/app/shapes"
	"github.com/ShaunBillows/shapes-cli-project-go/internal/messages"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   shapes.Shape
		hasArea float64
		err     error
	}{
		{name: "circle test: positive radius", shape: shapes.Circle{Radius: 10}, hasArea: 314.1592653589793, err: nil},
		{name: "circle test: negative radius", shape: shapes.Circle{Radius: -10}, hasArea: 0, err: messages.ErrNegativeRadius},
		{name: "rectangle test: positive width and height", shape: shapes.Rectangle{Height: 10, Width: 10}, hasArea: 100, err: nil},
		{name: "rectangle test: negative width", shape: shapes.Rectangle{Height: 10, Width: -10}, hasArea: 0, err: messages.ErrNegativeWidth},
		{name: "rectangle test: negative height", shape: shapes.Rectangle{Height: -10, Width: 10}, hasArea: 0, err: messages.ErrNegativeHeight},
		{name: "rectangle test: negative width and height", shape: shapes.Rectangle{Height: -10, Width: -10}, hasArea: 0, err: messages.ErrNegativeHeight},
		{name: "triangle test: positive base and height", shape: shapes.Triangle{Height: 10, Base: 10}, hasArea: 50, err: nil},
		{name: "triangle test: negative base", shape: shapes.Triangle{Height: 10, Base: -10}, hasArea: 0, err: messages.ErrNegativeBase},
		{name: "triangle test: negative base and height", shape: shapes.Triangle{Height: -10, Base: -10}, hasArea: 0, err: messages.ErrNegativeHeight},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shape.Area()
			assertError(t, tt.name, err, tt.err)
			assertFloat64(t, tt.name, got, tt.hasArea)
		})
	}
}
