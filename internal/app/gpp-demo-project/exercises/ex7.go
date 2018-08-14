package exercises

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type ExerciseSeven struct{}

// Go waiting groups
func (ex ExerciseSeven) Run() {
	wg = sync.WaitGroup{}

	fmt.Println("waiting group")
	wg.Add(2) // wait for two things

	go ex.test1()
	fmt.Println("ran test1()")
	go ex.test2()
	fmt.Println("ran test2()")

	wg.Wait() // wait for all things to finish executing

	fmt.Println("finished waiting")
}

func (ex ExerciseSeven) test1() {
	time.Sleep(5 * time.Second)
	fmt.Println("finished executing test1")
	wg.Done()
}

func (ex ExerciseSeven) test2() {
	fmt.Println("finished executing test2")
	wg.Done()
}
