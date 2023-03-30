package prompter

import (
	"errors"
	"github.com/manifoldco/promptui"
)

type PromptuiPrompts interface {
	Select(question string, options []string) (int, string, error)
	Prompt(data string) (string, error)
}

type PromptuiReader struct {
	Reader PromptuiPrompts
}

func (p *PromptuiReader) Prompt(data string) (string, error) {
	prompt := promptui.Prompt{
		Label: data,
	}
	return prompt.Run()
}

func (p *PromptuiReader) Select(question string, options []string) (int, string, error) {
	prompt := promptui.Select{
		Label: question,
		Items: options,
	}
	return prompt.Run()
}

func NewPromptuiPrompter() *PromptuiPrompter {
	reader := &PromptuiReader{}
	return &PromptuiPrompter{
		Reader: reader,
	}
}

type PromptuiPrompter struct {
	Reader PromptuiPrompts
}

func (p *PromptuiPrompter) GetUserChoice(question string, options []string) (string, error) {
	_, result, err := p.Reader.Select(question, options)
	if err != nil {
		return "", errors.New(ErrInvalidInput)
	}
	return result, nil
}

func (p *PromptuiPrompter) GetUserData(data string) (string, error) {
	result, err := p.Reader.Prompt(data)
	if err != nil {
		return "", errors.New(ErrInvalidInput)
	}
	return result, nil
}
