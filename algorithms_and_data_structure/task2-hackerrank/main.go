package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */
//this code

// ar []int32 — входной массив.
// sum — переменная для накопления суммы.
// for _, num := range ar — цикл по массиву.
// sum += num — прибавляем каждый элемент.


func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	for _, num := range ar {
		sum += num
	}
	return sum
}

// В Go, выражение for _, num := range ar позволяет перебрать все элементы массива, где:
// _ — это индекс элемента в массиве (он нам не нужен в данной задаче, поэтому мы его игнорируем, используя _).
// num — это сам элемент массива, который мы хотим обработать.
// Инициализируется sum = 0.
// Цикл по массиву ar = [1, 2, 3, 4, 10, 11]:
// Первый элемент 1: sum = 0 + 1 = 1
// Второй элемент 2: sum = 1 + 2 = 3
// Третий элемент 3: sum = 3 + 3 = 6
// Четвёртый элемент 4: sum = 6 + 4 = 10
// Пятый элемент 10: sum = 10 + 10 = 20
// Шестой элемент 11: sum = 20 + 11 = 31
// Финальная сумма: 31.
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := simpleArraySum(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
