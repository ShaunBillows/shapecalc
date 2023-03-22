package shapes

type ShapeType string

const (
	ShapeTypeRectangle ShapeType = "Rectangle"
	ShapeTypeTriangle  ShapeType = "Triangle"
	ShapeTypeCircle    ShapeType = "Circle"

	ErrNegativeRadius = "radius cannot be negative"
	ErrNegativeWidth  = "width cannot be negative"
	ErrNegativeHeight = "height cannot be negative"
	ErrNegativeBase   = "base cannot be negative"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
	Type() ShapeType
}
