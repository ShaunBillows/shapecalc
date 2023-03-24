package app

import (
	"errors"
	"github.com/ShaunBillows/shapes-cli-project-go/app/shapes"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			name:     "should return struct fields",
			input:    shapes.Rectangle{},
			expected: []string{"Height", "Width"},
		},
		{
			name:     "empty struct should return nil",
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
