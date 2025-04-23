package main

import "fmt"

func main() {
	maxi := 18
	a := []int{5, 10, 1, 2, 1, 20}
	count := 0
	sumi := 0
	//1) вам необходимо отсортировать массив по возрастанию
	//2) пройтись циклом и складывать все числа пока сумма этих цифр не будет больше maxi и потом сделать break
	//3) считать количество суммы
	for i := 0; i < len(a); i++ {
		//цикл для сравнения элементов массива
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for _, v := range a {
		if sumi+v > maxi {
			break
		}
		sumi += v
		count += 1
	}
	fmt.Println(count)
}
