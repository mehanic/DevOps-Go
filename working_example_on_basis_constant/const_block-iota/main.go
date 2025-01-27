package main

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Partyday
		numberOfDays // this constant is not exported
	)
	println(numberOfDays)
}

// 7
