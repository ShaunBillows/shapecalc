package app

import (
	"errors"
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
		name          string
		readerInput   string
		readerError   error
		expected      shapes.ShapeType
		expectedError error
	}{
		{
			name:        "option 1 should return a rectangle",
			readerInput: "1",
			expected:    shapes.ShapeTypeRectangle,
		},
		{
			name:        "option 2 should return a circle",
			readerInput: "2",
			expected:    shapes.ShapeTypeCircle,
		},
		{
			name:        "option 3 should return a triangle",
			readerInput: "3",
			expected:    shapes.ShapeTypeTriangle,
		},
		{
			name:          "should handle errors from stringReader",
			readerInput:   "4",
			readerError:   errors.New("Invalid readerInput. Please try again."),
			expectedError: errors.New("An error occurred while reading input"),
		},
		{
			name:          "should return an error with incorrect readerInput",
			readerInput:   "incorrect readerInput",
			expectedError: errors.New(ErrInvalidInput),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr.ReadStringFunc = func(delim byte) (string, error) {
				return tt.readerInput, tt.readerError
			}
			selectedShape, err := app.SelectShape()
			if err != nil {
				assertEquals(t, tt.expectedError.Error(), err.Error())
			}
			if err == nil {
				assertEquals(t, tt.expected, selectedShape.Type())
			}
		})
	}
}

func assertNotNil(t testing.TB, got interface{}) {
	t.Helper()
	if got == nil {
		t.Errorf("expected nil got %q", got)
	}
}

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
		t.Errorf("\nexpected `%v` \ngot `%v`", expected, got)
	}
}
