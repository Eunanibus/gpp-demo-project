package exercises

import (
	"fmt"
	"time"
)

type ExerciseSix struct {
}

// Go Routines
// A goroutine is a lightweight thread managed by the Go runtime

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func (ex ExerciseSix) Run() {
	// The `go say("X")` line spins a new goroutine running say("X")
	go say("GPP")
	say("Hello")
}
