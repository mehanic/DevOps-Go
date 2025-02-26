package main
import"fmt"
func main() {
	
	greeting := getGreeting
	result := greeting("Prasetiyo")
	fmt.Println(result)
	
}

func getGreeting(name string) string {
	return "Hallo " + name
}
