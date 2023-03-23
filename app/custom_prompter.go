package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CustomPrompter struct {
	Reader StringReader
}

func NewCustomPrompter() *CustomPrompter {
	reader := bufio.NewReader(os.Stdin)
	return &CustomPrompter{
		Reader: reader,
	}
}

func (p *CustomPrompter) GetUserChoice(question string, options []string) (string, error) {
	fmt.Print(question + "\n\n")
	var index int
	for i, v := range options {
		index = i + 1
		fmt.Printf("%v. %v\n", index, v)
	}
	fmt.Print("\nChoice : ")
	userInput, err := p.Reader.ReadString('\n')
	if err != nil {
		return "", errors.New(ErrReadingInput)
	}
	userInput = strings.TrimRight(userInput, "\n")
	var indexStr string
	for i, option := range options {
		index = i + 1
		indexStr = strconv.Itoa(index)
		if indexStr == userInput {
			return option, nil
		}
	}
	return "", errors.New(ErrInvalidInput)
}

func (p *CustomPrompter) GetUserData(data string) (string, error) {
	fmt.Printf("%v : ", data)
	userInput, err := p.Reader.ReadString('\n')
	if err != nil {
		return "", errors.New(ErrReadingInput)
	}
	userInput = strings.TrimRight(userInput, "\n")
	return userInput, nil
}
