package app

import (
	"errors"
	"github.com/ShaunBillows/shapes-cli-project-go/app/shapes"
	"github.com/stretchr/testify/assert"
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
			readerError:   errors.New("Invalid reader input."),
			expectedError: errors.New(ErrReadingInput),
		},
		{
			name:          "should return an error with incorrect readerInput",
			readerInput:   "incorrect reader input",
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

func TestInputReader(t *testing.T) {
	assert := assert.New(t)
	mr := &mockReader{}
	app := NewApp()
	app.Reader = mr
	tests := []struct {
		name          string
		readerInput   string
		readerOptions []string
		readerError   error
		expected      string
		expectedError error
	}{
		{
			name:          "should return the selected option",
			readerInput:   "1",
			readerOptions: []string{"1", "2", "3"},
			readerError:   nil,
			expected:      "1",
			expectedError: nil,
		},
		{
			name:          "should return an error if the input wasn't an option",
			readerInput:   "100",
			readerOptions: []string{"1", "2", "3"},
			readerError:   nil,
			expected:      "",
			expectedError: errors.New(ErrInvalidInput),
		},
		{
			name:          "should handle errors from stringReader",
			readerInput:   "1",
			readerOptions: []string{"1", "2", "3"},
			readerError:   errors.New("Invalid reader input."),
			expected:      "1",
			expectedError: errors.New(ErrReadingInput),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr.ReadStringFunc = func(delim byte) (string, error) {
				return tt.readerInput, tt.readerError
			}
			got, err := app.InputReader(tt.readerInput, tt.readerOptions)
			if err != nil {
				assert.Equal(tt.expectedError.Error(), err.Error(), tt.name)
			}
			if err == nil {
				assert.Equal(tt.expected, got, tt.name)
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
		t.Errorf("\nexpected `%v` \ngot `%v`", expected, got)
	}
}
