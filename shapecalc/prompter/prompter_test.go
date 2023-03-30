package prompter

import (
	"errors"
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
			name:          "should handle invalid input",
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
