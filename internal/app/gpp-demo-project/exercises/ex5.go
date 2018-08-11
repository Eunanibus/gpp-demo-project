package exercises

import "fmt"

<<<<<<< HEAD
// Interfaces

type ExerciseFive struct {
}

=======
type ExerciseFive struct {
}

// Interfaces

>>>>>>> e0e15892c921234ae4d367d99596e22491b28de5
type I interface {
	Test()
}

type T struct {
	message string
}

// This method means type T implements the interface I
func (s T) Test() {
	fmt.Println(s.message)
}

func (ex ExerciseFive) Run() {
	// Type T now implements the Test() function
	// It now conforms to the interface, though we don't need to explicitly declare it does so
	var i I = T{"This is a message"}
	i.Test()
	// Static type checking will do the rest
	// Try commenting out the Test() function and see what happens
}