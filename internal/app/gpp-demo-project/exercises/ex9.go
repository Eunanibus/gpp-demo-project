package exercises

import "fmt"

type ExerciseNine struct{}

type ExampleObject struct {
	name string
}

// Pointers and references
func (ex ExerciseNine) Run() {

	exOb := ExampleObject{
		name: "Test object",
	}

	// When passing an object to a function, the parameter you pass is actually a new object with a new reference in memory
	// In order to use the same object, you need to pass a pointer to the object

	// This will work because it is a pointer to the object
	objectPrinter(&exOb)

	// This will not work because it is a direct reference to the object
	//objectPrinter(exOb)
}

// Expected Pointers can be declared in signatures with the * notation
func objectPrinter(object *ExampleObject) {
	fmt.Println(object.name)
}
