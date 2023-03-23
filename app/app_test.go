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

func TestCustomPrompter_GetUserChoice(t *testing.T) {
	assert := assert.New(t)
	mr := &mockReader{}
	prompter := NewCustomPrompter()
	prompter.Reader = mr
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
			readerOptions: []string{"Rectangle", "Circle", "Triangle"},
			readerError:   nil,
			expected:      "Rectangle",
			expectedError: nil,
		},
		{
			name:          "invalid option should return an error",
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
			got, err := prompter.GetUserChoice(tt.readerInput, tt.readerOptions)
			if err != nil {
				assert.Equal(tt.expectedError.Error(), err.Error(), tt.name)
			}
			if err == nil {
				assert.Equal(tt.expected, got, tt.name)
			}
		})
	}
}

func TestCustomPrompter_GetUserData(t *testing.T) {
	assert := assert.New(t)
	mr := &mockReader{}
	prompter := NewCustomPrompter()
	prompter.Reader = mr
	tests := []struct {
		name          string
		readerInput   string
		readerData    string
		readerError   error
		expected      string
		expectedError error
	}{
		{
			name:          "should return the user's input",
			readerInput:   "user's input",
			readerError:   nil,
			expected:      "user's input",
			expectedError: nil,
		},
		{
			name:          "should handle errors from stringReader",
			readerInput:   "1",
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
			got, err := prompter.GetUserData(tt.readerData)
			if err != nil {
				assert.Equal(tt.expectedError.Error(), err.Error(), tt.name)
			}
			if err == nil {
				assert.Equal(tt.expected, got, tt.name)
			}
		})
	}
}

type MockPromptuiReader struct {
	SelectFunc func(question string, options []string) (int, string, error)
	PromptFunc func(data string) (string, error)
}

func (p *MockPromptuiReader) Prompt(data string) (string, error) {
	return p.PromptFunc(data)
}

func (p *MockPromptuiReader) Select(question string, options []string) (int, string, error) {
	return p.SelectFunc(question, options)
}

func TestPromptuiPrompter_GetUserChoice(t *testing.T) {
	assert := assert.New(t)
	mr := &MockPromptuiReader{}
	prompter := NewPromptuiPrompter()
	prompter.Reader = mr
	tests := []struct {
		name          string
		readerInput   string
		readerOptions []string
		expected      string
		expectedError error
	}{
		{
			name:          "should return the user's input",
			readerInput:   "1",
			readerOptions: []string{"Red", "Green", "Blue"},
			expected:      "Red",
			expectedError: nil,
		},
		{
			name:          "should handle invalid input",
			readerInput:   "4",
			readerOptions: []string{"Red", "Green", "Blue"},
			expected:      "",
			expectedError: errors.New(ErrInvalidInput),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr.SelectFunc = func(question string, options []string) (int, string, error) {
				return 0, tt.expected, tt.expectedError
			}
			got, err := prompter.GetUserChoice(tt.readerInput, tt.readerOptions)
			if err != nil {
				assert.Equal(tt.expectedError.Error(), err.Error(), tt.name)
			}
			if err == nil {
				assert.Equal(tt.expected, got, tt.name)
			}
		})
	}
}

func TestPromptuiPrompter_GetUserData(t *testing.T) {
	assert := assert.New(t)
	mr := &MockPromptuiReader{}
	prompter := NewPromptuiPrompter()
	prompter.Reader = mr
	tests := []struct {
		name          string
		readerInput   string
		expected      string
		expectedError error
	}{
		{
			name:          "should return the user's input",
			readerInput:   "4",
			expected:      "4",
			expectedError: nil,
		},
		{
			name:          "should handle reader errors",
			readerInput:   "4",
			expected:      "",
			expectedError: errors.New(ErrInvalidInput),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr.PromptFunc = func(data string) (string, error) {
				return tt.expected, tt.expectedError
			}
			got, err := prompter.GetUserData(tt.readerInput)
			if err != nil {
				assert.Equal(tt.expectedError.Error(), err.Error(), tt.name)
			}
			if err == nil {
				assert.Equal(tt.expected, got, tt.name)
			}
		})
	}
}

func TestApp_CreateShape(t *testing.T) {
	assert := assert.New(t)
	app := NewApp("custom")
	tests := []struct {
		name          string
		shape         string
		expected      shapes.Shape
		expectedError error
	}{
		{
			name:          "should return a circle",
			shape:         "Circle",
			expected:      shapes.NewCircle(),
			expectedError: nil,
		},
		{
			name:          "should return a rectangle",
			shape:         "Rectangle",
			expected:      shapes.NewRectangle(),
			expectedError: nil,
		},
		{
			name:          "should return a triangle",
			shape:         "Triangle",
			expected:      shapes.NewTriangle(),
			expectedError: nil,
		},
		{
			name:          "should return an error",
			shape:         "Invalid shape",
			expected:      nil,
			expectedError: errors.New(ErrInvalidInput),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := app.CreateShape(tt.shape)
			assert.Equal(tt.expectedError, err, tt.name)
			assert.Equal(tt.expected, got, tt.name)
		})
	}
}

func TestApp_BuildShape(t *testing.T) {
	assert := assert.New(t)
	app := NewApp("custom")
	tests := []struct {
		name          string
		shape         shapes.Shape
		shapeData     ShapeData
		expected      shapes.Shape
		expectedError error
	}{
		{
			name:  "should return a rectangle",
			shape: shapes.NewRectangle(),
			shapeData: ShapeData{
				"Height": 2,
				"Width":  2,
			},
			expected: &shapes.Rectangle{
				2,
				2,
			},
			expectedError: nil,
		},
		{
			name:  "should return a circle",
			shape: shapes.NewCircle(),
			shapeData: ShapeData{
				"Radius": 2,
			},
			expected: &shapes.Circle{
				2,
			},
			expectedError: nil,
		},
		{
			name:  "should return a triangle",
			shape: shapes.NewTriangle(),
			shapeData: ShapeData{
				"Height": 2,
				"Base":   2,
			},
			expected: &shapes.Triangle{
				2,
				2,
			},
			expectedError: nil,
		},
		{
			name:          "should handle invalid values",
			shape:         nil,
			shapeData:     nil,
			expected:      nil,
			expectedError: errors.New(ErrInvalidInput),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := app.BuildShape(tt.shape, tt.shapeData)
			assert.Equal(tt.expectedError, err, tt.name)
			assert.Equal(tt.expected, got, tt.name)
		})
	}
}

func TestApp_PerformShapeAction(t *testing.T) {
	assert := assert.New(t)
	app := NewApp("custom")
	tests := []struct {
		name          string
		shape         shapes.Shape
		action        string
		expected      float64
		expectedError error
	}{
		{
			name: "should return area",
			shape: shapes.Rectangle{
				2,
				2,
			},
			action:        "Area",
			expected:      4,
			expectedError: nil,
		},
		{
			name: "should return perimeter",
			shape: shapes.Rectangle{
				1,
				1,
			},
			action:        "Perimeter",
			expected:      4,
			expectedError: nil,
		},
		{
			name: "should handle invalid input",
			shape: shapes.Rectangle{
				2,
				2,
			},
			action:        "Invalid action",
			expected:      0,
			expectedError: errors.New(ErrInvalidInput),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := app.PerformShapeAction(tt.shape, tt.action)
			assert.Equal(tt.expectedError, err, tt.name)
			assert.Equal(tt.expected, got, tt.name)
		})
	}
}

func TestApp_GetFields(t *testing.T) {
	assert := assert.New(t)
	app := NewApp("custom")
	tests := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{
			name:     "test with struct",
			input:    shapes.Rectangle{},
			expected: []string{"Height", "Width"},
		},
		{
			name:     "test with empty struct",
			input:    struct{}{},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := app.GetFields(tt.input)
			assert.Equal(tt.expected, result, tt.name)
		})
	}
}
