package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ShaunBillows/shapes-cli-project-go/app/shapes"
	_struct "github.com/ShaunBillows/shapes-cli-project-go/internal/struct"
	"os"
	"strconv"
	"strings"
)

const (
	ErrInvalidInput = "Invalid input. Please try again."
)

type App struct {
	Reader StringReader
}

func NewApp() *App {
	reader := bufio.NewReader(os.Stdin)
	return &App{
		Reader: reader,
	}
}

func (a *App) Run() {

	var shapeSelected shapes.Shape
	var err error
	for {
		shapeSelected, err = a.SelectShape()
		if shapeSelected != nil {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	actionSelected := a.selectShapeAction(shapeSelected)

	shapeDimensions := a.selectDimensions(shapeSelected)

	s := a.buildShape(shapeSelected, shapeDimensions)

	r, _ := a.calculateResult(s, actionSelected)

	fmt.Printf("\n\nThe %v of the %T is %v\n\n", actionSelected, shapeSelected, r)
}

func (a *App) SelectShape() (shapes.Shape, error) {

	fmt.Print("Select a shape (enter 1,2 or 3):\n1. Rectangle\n2. Circle\n3. Triangle\nChoice : ")
	userInput, err := a.Reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("An error occurred while reading input")
	}
	shapeSelected := strings.TrimRight(userInput, "\n")
	switch shapeSelected {
	case "1":
		return shapes.NewRectangle(), nil
	case "2":
		return shapes.NewCircle(), nil
	case "3":
		return shapes.NewTriangle(), nil
	default:
		fmt.Println("Invalid input. Please try again.")
		return nil, errors.New(ErrInvalidInput)
	}
}

func (a *App) selectShapeAction(s shapes.Shape) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which operation would you like to perform? (enter 1 or 2):\n1. Calculate area\n2. Calculate perimeter\nChoice : ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		os.Exit(1)
	}
	actionSelected := strings.TrimRight(userInput, "\n")
	switch actionSelected {
	case "1":
		return "Area"
	case "2":
		return "Perimeter"
	default:
		fmt.Println("Invalid input. Please try again.")
		return a.selectShapeAction(s)
	}
}

func (a *App) selectDimensions(s shapes.Shape) shapeData {
	reader := bufio.NewReader(os.Stdin)
	fields := _struct.GetFields(s)
	var userInput string
	var userInputVal float64
	var err error
	dimensions := map[string]float64{}
	for _, param := range fields {
		fmt.Printf("Enter %v : ", param)
		userInput, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.")
			os.Exit(1)
		}
		userInput = strings.TrimRight(userInput, "\n")
		userInputVal, err = strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("You must enter a number. Please try again.")
			return a.selectDimensions(s)
		}
		dimensions[param] = userInputVal
	}
	return dimensions
}

func (a *App) buildShape(s shapes.Shape, d shapeData) shapes.Shape {
	switch s.(type) {
	case *shapes.Rectangle:
		r := s.(*shapes.Rectangle)
		r.Width = d["Width"]
		r.Height = d["Height"]
		return r
	case *shapes.Circle:
		c := s.(*shapes.Circle)
		c.Radius = d["Radius"]
		return c
	case *shapes.Triangle:
		t := s.(*shapes.Triangle)
		t.Height = d["Height"]
		t.Base = d["Base"]
		return t
	default:
		fmt.Println("Invalid shape type.")
		return nil
	}
}

func (a *App) calculateResult(s shapes.Shape, action string) (float64, error) {
	switch action {
	case "Area":
		result, err := s.Area()
		if err != nil {
			return 0, err
		}
		return result, nil
	case "Perimeter":
		result, err := s.Perimeter()
		if err != nil {
			return 0, err
		}
		return result, nil
	default:
		return 0, nil
	}
}

type shapeData map[string]float64
