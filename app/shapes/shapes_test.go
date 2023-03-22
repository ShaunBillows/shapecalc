package shapes

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
		err          error
	}{
		{name: "circle test: positive radius", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586, err: nil},
		{name: "circle test: negative radius", shape: Circle{Radius: -10}, hasPerimeter: 0, err: errors.New(ErrNegativeRadius)},
		{name: "rectangle test: positive width and height", shape: Rectangle{Height: 10, Width: 10}, hasPerimeter: 40, err: nil},
		{name: "rectangle test: negative width", shape: Rectangle{Height: 10, Width: -10}, hasPerimeter: 0, err: errors.New(ErrNegativeWidth)},
		{name: "rectangle test: negative height", shape: Rectangle{Height: -10, Width: 10}, hasPerimeter: 0, err: errors.New(ErrNegativeHeight)},
		{name: "rectangle test: negative width and height", shape: Rectangle{Height: -10, Width: -10}, hasPerimeter: 0, err: errors.New(ErrNegativeHeight)},
		{name: "triangle test: positive base and height", shape: Triangle{Height: 3, Base: 5}, hasPerimeter: 12.810249675906654, err: nil},
		{name: "triangle test: negative base", shape: Triangle{Height: 10, Base: -10}, hasPerimeter: 0, err: errors.New(ErrNegativeBase)},
		{name: "triangle test: negative base and height", shape: Triangle{Height: -10, Base: -10}, hasPerimeter: 0, err: errors.New(ErrNegativeHeight)},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shape.Perimeter()
			assert.Equal(t, err, tt.err, tt.name)
			assertFloat64(t, tt.name, got, tt.hasPerimeter)
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
		err     error
	}{
		{name: "circle test: positive radius", shape: Circle{Radius: 10}, hasArea: 314.1592653589793, err: nil},
		{name: "circle test: negative radius", shape: Circle{Radius: -10}, hasArea: 0, err: errors.New(ErrNegativeRadius)},
		{name: "rectangle test: positive width and height", shape: Rectangle{Height: 10, Width: 10}, hasArea: 100, err: nil},
		{name: "rectangle test: negative width", shape: Rectangle{Height: 10, Width: -10}, hasArea: 0, err: errors.New(ErrNegativeWidth)},
		{name: "rectangle test: negative height", shape: Rectangle{Height: -10, Width: 10}, hasArea: 0, err: errors.New(ErrNegativeHeight)},
		{name: "rectangle test: negative width and height", shape: Rectangle{Height: -10, Width: -10}, hasArea: 0, err: errors.New(ErrNegativeHeight)},
		{name: "triangle test: positive base and height", shape: Triangle{Height: 10, Base: 10}, hasArea: 50, err: nil},
		{name: "triangle test: negative base", shape: Triangle{Height: 10, Base: -10}, hasArea: 0, err: errors.New(ErrNegativeBase)},
		{name: "triangle test: negative base and height", shape: Triangle{Height: -10, Base: -10}, hasArea: 0, err: errors.New(ErrNegativeHeight)},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shape.Area()
			assert.Equal(t, err, tt.err, tt.name)
			assert.Equal(t, got, tt.hasArea, tt.name)
		})
	}
}

func TestTypes(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasType ShapeType
	}{
		{name: "circle should return circle type`", shape: Circle{Radius: 10}, hasType: ShapeTypeCircle},
		{name: "rectangle should return rectangle type`", shape: Rectangle{Height: 10, Width: 10}, hasType: ShapeTypeRectangle},
		{name: "triangle should return triangle type`", shape: Triangle{Height: 10, Base: 10}, hasType: ShapeTypeTriangle},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Type()
			assert.Equal(t, got, tt.hasType, tt.name)
		})
	}
}

func assertFloat64(t testing.TB, name string, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %g want %g", name, got, want)
	}
}
