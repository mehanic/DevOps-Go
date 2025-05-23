package main

func main() {
	// Switch - Case
	var switchVar any
	switch switchVar {
	case "1":
		println("Hello world")
		// break -> automatic
		fallthrough // continue validating next case
	case "2", "3":
		fallthrough
	case "5":
		return
	default:
		println("Default actions")
	}

	// Struct comparison in switch
	var structValue struct{ ID uint64 }
	switch structValue {
	case struct{ ID uint64 }{
		ID: 0,
	}:
		println("Matches")
	}

	// Switch - Type assertion
	var myVariable any = "string type"
	switch myValue := myVariable.(type) {
	case int:
		println(myValue)
	case uint:
		println(myValue)
	case float32:
		println(myValue)
	case float64:
		println(myValue)
	case string:
		println(myValue)
	default:
		println("Unknown type")
	}
}// Default actions
// Matches
// string type