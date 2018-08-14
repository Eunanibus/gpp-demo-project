package exercises

import "fmt"

type ExerciseEight struct{}

// Go waiting groups
func (ex ExerciseEight) Run() {
	ex.test(myCoolFunction)
}

// You can also pass functions
func (ex ExerciseEight) test(myFunc func(int) string) {
	fmt.Println("executing test")

	myFuncResult := myFunc(1337)
	fmt.Printf("myFunc() output - %s", myFuncResult)

	fmt.Println("finished executing test")
}

func myCoolFunction(i int) string {
	return fmt.Sprintf("%d\n", i)
}
