// Изучаем Golang. Урок 3. Основы. Продолжение. Указатели, структуры, массивы и слайсы

package main

import (
	"fmt"
)

type Point struct { // Тип Структуры
	X int    // Поля структуры
	Y int    // Поля структуры
	S string // Поля структуры
}

func (p Point) method() { // Методы Структуры
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

//func (p *Point) method() { // Методы Структуры
//	fmt.Println(p.X)
//	fmt.Println(p.Y)
//	fmt.Println(p.S)
//}

func main() {
	// Указатели
	a := "Указатели"       // Инициализируем переименую
	p := &a                // Получить ссылку
	fmt.Println("p =", p)  // Получить адрес ссылки
	fmt.Println("p =", *p) // Получить значения по ссылке
	*p = "Указатели!"      // Изменить значения по ссылке
	fmt.Println("a =", a)

	b := 42     // Инициализируем переименую (поля)
	g := &b     // Получить ссылку
	*g = *g / 2 // Изменить значения по ссылке
	fmt.Println("b =", b)

	// Структуры
	structs() // Вызываем Функцию structs

	p3 := Point{ // Заполняем Структуру №3
		X: 10,
		Y: 20,
		S: "Строка?",
	}
	p3.method() // Вызываем Функцию method
	p4 := &p3
	p4.method() // Вызываем Функцию method

	// Массивы
	var strArray [3]string
	strArray[0] = "first"
	strArray[1] = "second"
	fmt.Println("strArray: ", strArray)
	fmt.Println("strArray: ", strArray[1])

	intArray := [...]int{1, 2, 3} // Инициализация Массива и авто размерность
	fmt.Println("intArray: ", intArray)

	// Слайсы
	letters := []string{"a", "b", "c"} // Инициализация Слайс
	letters[0] = "a1"                  // Меняем значения элемента в letters
	letters = append(letters, "d", "e", "f")
	fmt.Println("Слайс letters ", letters)
	fmt.Println("Слайс letters ", letters[5])

	createSlice := make([]string, 3) // Инициализация пустой Слайс
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	fmt.Println(fmt.Sprintf("len: %d", len(createSlice))) // len
	fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice))) // cap
	createSlice = append(createSlice, "4")
	fmt.Println("Слайс createSlice ", createSlice)

	animalsArr := [4]string{ // array
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}
	animalsSlic := []string{ // slices
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	var arr1 []string = animalsArr[0:2]
	arr2 := animalsArr[1:3]
	fmt.Println("animalsArr ", arr1)
	fmt.Println("animalsArr ", arr2)
	var slic1 []string = animalsSlic[0:2]
	slic2 := animalsSlic[1:3]
	fmt.Println("animalsSlic ", slic1)
	fmt.Println("animalsSlic ", slic2)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := s[0:5] //Нарезка массива c 0 до 5 элемента
	t3 := s[:5] //Нарезка массива c 0 до 5 элемента (упрошенный вариант)
	fmt.Println("Нарезка массива c 0 до 5 элемента: ", t)
	fmt.Println("Нарезка массива c 0 до 5 элемента: ", t3)
	t2 := s[5:10] //Нарезка массива c 5 до 10 элемента
	t4 := s[5:]   //Нарезка массива c 5 до 10 элемента (упрошенный вариант)
	fmt.Println("Нарезка массива c 5 до 10  элемента: ", t2)
	fmt.Println("Нарезка массива c 5 до 10  элемента: ", t4)
	t5 := s[:] //Нарезка массива c 0 до 10 элемента (упрошенный вариант)
	fmt.Println("Нарезка массива c 0 до 10  элемента: ", t5)

	// Пример пустого Слайса
	var snil []int
	fmt.Println("Пустой Слайс: ", snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("slice is nil!")
	}
}

// Функций Структуры
func structs() {
	p1 := Point{ // Заполняем Структуру №1
		X: 1,
		Y: 2,
		S: "Строка",
	}
	fmt.Println("Point_1", p1)
	fmt.Println("p1.X", p1.X)
	fmt.Println("p1.Y", p1.Y)
	fmt.Println("p1.S", p1.S)
	p1.X = 3
	fmt.Println("p1.X", p1.X)

	p2 := Point{X: 4} // Заполняем Структуру №2
	fmt.Println("Point_2", p2)

	j := &p1
	fmt.Println("Point_2", &j) // Получить адрес ссылки из Структуры
	fmt.Println("Point_2", j)  // Получить ссылку из Структуры
	fmt.Println("Point_2", *j) // Получить значения по ссылке из Структуры
	fmt.Println("p2.S", j.S)   // Получить значения по ссылке из Структуры
	fmt.Println(*&j)           // Создаем ссылку, потом ее разыменовываем
}
