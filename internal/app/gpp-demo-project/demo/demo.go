package demo

import "github.com/Eunanibus/gpp-demo-project/internal/app/gpp-demo-project/exercises"

func RunDemo() {
	runExercise(exercises.ExerciseOne{})
	runExercise(exercises.ExerciseTwo{})
	runExercise(exercises.ExerciseThree{})
	runExercise(exercises.ExerciseFour{})
	runExercise(exercises.ExerciseFive{})
	runExercise(exercises.ExerciseSix{})
}

func runExercise(exercise exercises.Exercise) {
	exercise.Run()
}