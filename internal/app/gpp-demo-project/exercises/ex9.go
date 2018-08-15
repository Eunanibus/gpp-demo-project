package exercises

import "fmt"

type ExerciseNine struct{}

type ExampleObject struct {
	name string
}

// Pointers and references
func (ex ExerciseNine) Run() {

	exOb1 := ExampleObject{
		name: "Test object",
	}

	exOb2 := ExampleObject{
		name: "Test object",
	}

	mutateObjectPointer(&exOb1)
	fmt.Printf("exOb1 name - %s ", exOb1.name)

	mutateObject(exOb2)
	fmt.Printf("exOb2 name - %s ", exOb2.name)

	exOb2 = mutateObjectWithReturn(exOb2)
	fmt.Printf("exOb2 name - %s ", exOb2.name)
}

func mutateObjectPointer(object *ExampleObject) {
	object.name = "Cats are cool"
}

func mutateObject(object ExampleObject) {
	object.name = "Cats are cool"
}

func mutateObjectWithReturn(object ExampleObject) ExampleObject {
	object.name = "Cats are cool"
	return object
}