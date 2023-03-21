package app

import (
	"errors"
	"fmt"
)

func (a *App) InputReader(question string, options []string) (string, error) {
	fmt.Println(question)
	var index int
	for i, v := range options {
		index = i + 1
		fmt.Printf("%v. %v\n", index, v)
	}
	userInput, err := a.Reader.ReadString('\n')
	if err != nil {
		return "", errors.New(ErrReadingInput)
	}
	for _, option := range options {
		if userInput == option {
			fmt.Println(userInput)
			return userInput, nil
		}
	}
	return "", errors.New(ErrInvalidInput)
}
