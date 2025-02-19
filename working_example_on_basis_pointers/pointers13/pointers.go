package main

func main() {

	u1, u2 := "orket", "Paul"
	p := &u1
	*p = u2
	println(u1, u2)

}
