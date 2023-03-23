package app

import (
	"github.com/manifoldco/promptui"
)

type PromptuiSelect struct{}

func (p *PromptuiSelect) Select(question string, options []string) (int, string, error) {
	prompt := promptui.Select{
		Label: question,
		Items: options,
	}
	return prompt.Run()
}

type SelectPrompter interface {
	Select(question string, options []string) (int, string, error)
}

type DataPrompter interface {
	Prompt(data string) (string, error)
}

type PromptuiPrompter struct {
	SelectPrompter SelectPrompter
	DataPrompter   DataPrompter
}

type PromptuiPrompt struct{}

func (p *PromptuiPrompt) Prompt(data string) (string, error) {
	prompt := promptui.Prompt{
		Label: data,
	}
	return prompt.Run()
}

func NewPromptuiPrompter() *PromptuiPrompter {
	selectPrompter := &PromptuiSelect{}
	dataPrompter := &PromptuiPrompt{}
	return &PromptuiPrompter{
		SelectPrompter: selectPrompter,
		DataPrompter:   dataPrompter,
	}
}

func (p *PromptuiPrompter) GetUserChoice(question string, options []string) (string, error) {
	_, result, err := p.SelectPrompter.Select(question, options)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (p *PromptuiPrompter) GetUserData(data string) (string, error) {
	result, err := p.DataPrompter.Prompt(data)
	if err != nil {
		return "", err
	}
	return result, nil
}

//////////////////////////////////////////////////////////////

//package app
//
//import (
//	"github.com/manifoldco/promptui"
//)
//
//type PromptuiPrompter struct {
//}
//
//func NewPromptuiPrompter() *PromptuiPrompter {
//
//	return &PromptuiPrompter{}
//}
//
//func (p *PromptuiPrompter) GetUserChoice(question string, options []string) (string, error) {
//	prompt := promptui.Select{
//		Label: question,
//		Items: options,
//	}
//
//	_, result, err := prompt.Run()
//
//	if err != nil {
//		return "", err
//	}
//
//	return result, nil
//}
//
//func (p *PromptuiPrompter) GetUserData(data string) (string, error) {
//	prompt := promptui.Prompt{
//		Label: data,
//	}
//
//	result, err := prompt.Run()
//
//	if err != nil {
//		return "", err
//	}
//
//	return result, nil
//}
