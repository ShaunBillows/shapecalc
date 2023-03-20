package messages

const (
	ErrNegativeRadius = ShapeErr("radius cannot be negative")
	ErrNegativeWidth  = ShapeErr("width cannot be negative")
	ErrNegativeHeight = ShapeErr("height cannot be negative")
	ErrNegativeBase   = ShapeErr("base cannot be negative")
)

type ShapeErr string

func (e ShapeErr) Error() string {
	return string(e)
}
