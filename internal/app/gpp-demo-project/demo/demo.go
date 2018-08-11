package demo

import "github.com/Eunanibus/gpp-demo-project/internal/app/gpp-demo-project/exercises"

func RunDemo() {
	runExercise(exercises.ExerciseOne{})
}

func runExercise(exercise exercises.Exercise) {
	exercise.Run()
}