package demo

import "github.com/Eunanibus/gpp-demo-project/internal/app/gpp-demo-project/exercises"

func RunDemo() {
	var exercise exercises.Exercise = exercises.ExerciseOne{}
	exercise.Run()
}

func runExercise(exercise exercises.Exercise) {
}