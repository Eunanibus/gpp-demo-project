package exercises

import "fmt"

// Multiple returns

type ExerciseThree struct {
}

func swap(x, y string) (string, string) {
	return y, x
}

func (ex ExerciseThree) Run() {
	a, b := swap("GPP!", "Hello, ")
	fmt.Println(a, b)
}
