package main
import "fmt"

func CreatePhoneNumber(numbers [10]uint) string {
	mapk := make(map[int]string)
	mapk[0] = "("
	mapk[3] = ") "
	mapk[6] = "-"

	result := ""
	for i, num := range numbers {
		if str, ok := mapk[i]; ok {
			result = result + str
		}
		result = fmt.Sprintf("%s%v", result, num)
	}
	return result
}

func main() {
	n := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(CreatePhoneNumber(n))

}
