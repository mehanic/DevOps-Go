package __6

import (
	"fmt"
	_ "fmt"
)

var record = struct {
	Name string
	Age  int
}{
	Name: "John Doak",
	Age:  100, // Yeah, not publishing the real one

}

func new() {
	fmt.Printf("%s is %d years old\n", record.Name, record.Age)
}

// -----------------
type Record struct {
	Name string
	Age  int
}

func main() {
	david := Record{Name: "David Justice", Age: 28}
	sarah := Record{Name: "Sarah Murphy", Age: 28}
	fmt.Printf("%+v\n", david)
	fmt.Printf("%+v\n", sarah)
}

// String returns a csv representing our record.
func (r Record) String() string {
	return fmt.Sprintf("%s,%d", r.Name, r.Age)

}

john := Record {Name: "John Doak", Age: 100}
fmt.Println(john.String())
//A method is similar to a function, but instead of being independent, it is bound to a
//type.

//Changing a field's value
myRecord.Name = "Peter Griffin"
fmt.Println(myRecord.Name) // Prints: Peter Griffin

func changeName(r Record) {
	r.Name = "Peter"
	fmt.Println("inside changeName: ", r.Name)
}
func main1() {
	rec := Record{Name: "John"}
	changeName(rec)
	fmt.Println("main: ", rec.Name)
}
//section on pointers, this is because the variable is copied, and we are
//changing the copy
func changeName1(r *Record) {
	r.Name = "Peter"
	fmt.Println("inside changeName: ", r.Name)
}
func main2() {
	// Create a pointer to a Record
	rec := &Record{Name: "John"}
	changeName1(rec)
	fmt.Println("main: ", rec.Name)
}


//Changing a field's value in a method
func (r Record) IncrAge() {
	r.Age++
}

func (r *Record) IncrAge1() {
	r.Age++
}