package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа
	var num2 = 2

	fmt.Println(num0)
	fmt.Println(num1)
	fmt.Println(num2)

	// короткое объявление переменной
	num3 := 3
	fmt.Println("Переменная num3 - ", num3)

	// Объявление нескольких переменных
	var weight, height int = 10, 20
	fmt.Println(weight, height)

	// Короткое присваивание
	// Хотя бы одна из переменных должна быть новой
	a, b := 14, 15
	fmt.Println("Переменные a, b - ", a, b)

	// Типы переменных
	// int - платформенно зависимый тип
	var i int = 10
	fmt.Println("i - ", i)
	// автоматически выбранный тип int
	var autoInt = -10
	fmt.Println("autoint - ", autoInt)
	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1
	fmt.Println("bigInt - ", bigInt)

	// числа с плавающей точкой float32, float64
	var floatNum float32 = 2.64
	fmt.Println("floatNum", floatNum)

	// булевый тип
	var booleanboo bool = true
	fmt.Println("boolean b - ", booleanboo)

	// Строки
	// по умолчанию пустая строка
	var str string
	fmt.Println("Строка", str, "Строка")
	// со спец символами
	var helloWorld string = "Привет\n\t"
	fmt.Println(helloWorld)
	// без спецсимволов
	var world string = `Привет мир\n\t`
	fmt.Println(world)

	// UTF-8 из коробки
	// Любые символы, хоть китайские

	// одинарные кавычки для байт (uint8)
	var rawBinary byte = '\x27'
	fmt.Println(rawBinary)

	// rune(uint32) для UTF-8 символов
	var someChinese rune = '临'
	fmt.Println(someChinese)

	// конкатенация
	hi := "Привет Мир!"
	andGoodMorning := hi + " И доброе утро!"
	fmt.Println(andGoodMorning)

	// строки неизменяемы
	// Нельзя hi[2] = 72

	// получение длины строки в байтах и символах
	byteLen := len(hi)
	symbols := utf8.RuneCountInString(hi)
	fmt.Println("Длина строки в байтах", byteLen)
	fmt.Println("Сколько символов в строке", symbols)

	// получение подстроки, в байтах не символах!
	hiSubstring := hi[:12] // Привет, 0-11 байты
	H := hi[0]             // byte, 72, не 'П'

	fmt.Println("получение подстроки, в байтах не символах!")
	fmt.Println("hiSubstring = ", hiSubstring)
	fmt.Println("H = ", H)

	// конвертация в слайс байт и обратно
	byteString := []byte(hi)
	afterbyteString := string(byteString)
	fmt.Println("byteString = ", byteString)
	fmt.Println("afterbyteString = ", afterbyteString)
}

// 0
// 1
// 2
// Переменная num3 -  3
// 10 20
// Переменные a, b -  14 15
// i -  10
// autoint -  -10
// bigInt -  4294967295
// floatNum 2.64
// boolean b -  true
// Строка  Строка
// Привет

// Привет мир\n\t
// 39
// 20020
// Привет Мир! И доброе утро!
// Длина строки в байтах 20
// Сколько символов в строке 11
// получение подстроки, в байтах не символах!
// hiSubstring =  Привет
// H =  208
// byteString =  [208 159 209 128 208 184 208 178 208 181 209 130 32 208 156 208 184 209 128 33]
// afterbyteString =  Привет Мир!
