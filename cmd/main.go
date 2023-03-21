package main

import (
	"bufio"
	"fmt"
	"github.com/ShaunBillows/shapes-cli-project-go/internal/shape"
	"github.com/ShaunBillows/shapes-cli-project-go/internal/shape/shapes"
	_struct "github.com/ShaunBillows/shapes-cli-project-go/internal/struct"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	shapeSelected := selectShape(reader)

	actionSelected := selectShapeAction(reader, shapeSelected)

	shapeDimensions := selectDimensions(reader, shapeSelected)

	s := buildShape(shapeSelected, shapeDimensions)

	r, _ := calculateResult(s, actionSelected)

	fmt.Printf("\n\nThe %v of the %T is %v\n\n", actionSelected, shapeSelected, r)

	// fmt.Printf("shape selected: %T\n", shapeSelected)
	// fmt.Printf("action selected %v\n", actionSelected)
	// fmt.Printf("shape dimensions %v\n", shapeDimensions)
}

func selectShape(reader *bufio.Reader) shape.Shape {
	fmt.Print("Select a shape (enter 1,2 or 3):\n1. Rectangle\n2. Circle\n3. Triangle\nChoice : ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again.")
		os.Exit(1)
	}
	shapeSelected := strings.TrimRight(userInput, "\n")
	switch shapeSelected {
	case "1":
		fmt.Println("You have selected a rectangle.")
		return shapes.NewRectangle()
	case "2":
		fmt.Println("You have selected a circle.")
		return shapes.NewCircle()
	case "3":
		fmt.Println("You have selected a triangle.")
		return shapes.NewTriangle()
	default:
		fmt.Println("Invalid input. Please try again.")
		return selectShape(reader)
	}
}

func selectShapeAction(reader *bufio.Reader, s shape.Shape) string {
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
		return selectShapeAction(reader, s)
	}
}

func selectDimensions(reader *bufio.Reader, s shape.Shape) shapeData {
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
			return selectDimensions(reader, s)
		}
		dimensions[param] = userInputVal
	}
	return dimensions
}
 
func buildShape(s shape.Shape, d shapeData) shape.Shape {
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

func calculateResult(s shape.Shape, a string) (float64, error) {
	switch a {
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
