package app

import (
	"github.com/ShaunBillows/shapes-cli-project-go/app/shapes"
	"testing"
)

type mockReader struct {
	ReadStringFunc func(delim byte) (string, error)
}

func (r *mockReader) ReadString(delim byte) (string, error) {
	return r.ReadStringFunc(delim)
}

func TestApp_SelectShape(t *testing.T) {
	mr := &mockReader{}

	app := NewApp()
	app.Reader = mr

	tests := []struct {
		name     string
		input    string
		expected shapes.ShapeType
		err      error
	}{
		{
			name:     "option 1 should return a rectangle",
			input:    "1",
			expected: shapes.ShapeTypeRectangle,
			err:      nil,
		},
		{
			name:     "option 2 should return a circle",
			input:    "2",
			expected: shapes.ShapeTypeCircle,
			err:      nil,
		},
		{
			name:     "option 3 should return a triangle",
			input:    "3",
			expected: shapes.ShapeTypeTriangle,
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr.ReadStringFunc = func(delim byte) (string, error) {
				return tt.input, tt.err
			}
			selectedShape, err := app.SelectShape()

			assertEquals(t, tt.err, err)
			if err == nil {
				assertEquals(t, tt.expected, selectedShape.Type())
			}
		})
	}
}

//func assertNotNil(t testing.TB, got interface{}) {
//	t.Helper()
//	if got == nil {
//		t.Errorf("expected nil got %q", got)
//	}
//}

//func assertError(t testing.TB, name string, expected, got error) {
//	t.Helper()
//	if got != expected {
//		t.Errorf("%#v expected %q got error %q", name, expected, got)
//	}
//}

//func assertFloat64(t testing.TB, name string, expected, got float64) {
//	t.Helper()
//	if got != expected {
//		t.Errorf("%#v  expected %g got %g", name, expected, got)
//	}
//}

func assertEquals(t testing.TB, expected interface{}, got interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("expected %v got %v", expected, got)
	}
}
