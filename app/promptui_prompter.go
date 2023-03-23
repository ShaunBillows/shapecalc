package app

import (
	"github.com/manifoldco/promptui"
)

type PromptuiPrompter struct {
}

func NewPromptuiPrompter() *PromptuiPrompter {
	return &PromptuiPrompter{}
}

func (p *PromptuiPrompter) GetUserChoice(question string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: question,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func (p *PromptuiPrompter) GetUserData(data string) (string, error) {
	prompt := promptui.Prompt{
		Label: data,
	}

	result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}
