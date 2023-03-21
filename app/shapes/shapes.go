package shapes

type ShapeType string

const (
	ShapeTypeRectangle ShapeType = "Rectangle"
	ShapeTypeTriangle  ShapeType = "Triangle"
	ShapeTypeCircle    ShapeType = "Circle"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
	Type() ShapeType
}
