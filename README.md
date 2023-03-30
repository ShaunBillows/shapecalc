# shapecalc

shapecalc is a simple Go package that provides functionality to calculate the area and perimeter of various shapes. It can be used as a standalone CLI application or as a package in your Go project.

## Installation

To install the package, simply run:

```bash
go get github.com/ShaunBillows/shapes-cli-project-go/shapecalc
```

## Usage

To use the package, simply import it into your Go project:

```go
import "github.com/ShaunBillows/shapes-cli-project-go/shapecalc"
```

Then, create a new instance of the App struct and call the Run() method:

```go
app := shapecalc.NewApp()
app.Run()
```

The Run() method will prompt the user for input and return the area and perimeter of the selected shape.

## Supported Shapes

Currently, the package supports the following shapes:

- Circle
- Rectangle
- Triangle
