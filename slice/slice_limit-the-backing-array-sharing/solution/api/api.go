package api

var temps = []int{5, 10, 3, 25, 45, 80, 90}

func Read(start, stop int) []int {
	portion := temps[start:stop:stop]

	return portion
}

func All() []int {
	return temps
}
