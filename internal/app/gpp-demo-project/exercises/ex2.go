package exercises

import "fmt"

// Structs
// These are typed collection of fields. Useful for grouping data together.

type ExerciseTwo struct {
}

type Rectangle struct {
	Name          string
	Width, Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

func (ex ExerciseTwo) Run() {
	rectangle := Rectangle{
		Name: "GPP Rectangle",
		Width: 20.5,
		Height: 40.5,
	}
	// Struct fields are accessed using a dot
	fmt.Printf("The height of my rectangle is %.2f" , rectangle.Width)
	fmt.Printf("The width of my rectangle is %.2f" , rectangle.Height)


	fmt.Printf("The area of my rectangle is %.2f" , rectangle.area())
}