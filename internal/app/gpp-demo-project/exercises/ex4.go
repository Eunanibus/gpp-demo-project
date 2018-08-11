package exercises

import "fmt"

// Pointers
// The & operator generates a pointer to its operand
// The * operator denotes the pointers underlying value

type ExerciseFour struct {
}

func (ex ExerciseFour) Run() {
	i, j := "I am variable i", "I am variable j"

	p := &i                      // point to i
	fmt.Println(*p)              // read i through the pointer
	*p = "I was once variable i" // set i through the pointer
	fmt.Println(i)               // see the new value of i

	p = &j                       // point to j
	*p = "I was once variable j" // reset j through the pointer
	fmt.Println(j)               // see the new value of j
}