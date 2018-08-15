package demo

import "github.com/eunanibus/gpp-demo-project/internal/app/gpp-demo-project/exercises"

func RunDemo() {
	runExercise(exercises.ExerciseOne{})
	runExercise(exercises.ExerciseTwo{})
	runExercise(exercises.ExerciseThree{})
	runExercise(exercises.ExerciseFour{})
	runExercise(exercises.ExerciseFive{})
	runExercise(exercises.ExerciseSix{})
	runExercise(exercises.ExerciseSeven{})
	runExercise(exercises.ExerciseEight{})
	runExercise(exercises.ExerciseNine{})
}

func runExercise(exercise exercises.Exercise) {
	exercise.Run()
}
