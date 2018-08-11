package demo

import "github.com/Eunanibus/gpp-demo-project/internal/app/gpp-demo-project/exercises"

func RunDemo() {
<<<<<<< HEAD
	runExercise(exercises.ExerciseOne{})
	runExercise(exercises.ExerciseTwo{})
	runExercise(exercises.ExerciseThree{})
	runExercise(exercises.ExerciseFour{})
	runExercise(exercises.ExerciseFive{})
=======
	//runExercise(exercises.ExerciseOne{})
>>>>>>> e0e15892c921234ae4d367d99596e22491b28de5
	runExercise(exercises.ExerciseSix{})
}

func runExercise(exercise exercises.Exercise) {
	exercise.Run()
}