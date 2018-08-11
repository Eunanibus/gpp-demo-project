package exercises

type ExerciseTwo struct {
}

type Rectangle struct {
	Name          string
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (ex ExerciseTwo) Run() {

}