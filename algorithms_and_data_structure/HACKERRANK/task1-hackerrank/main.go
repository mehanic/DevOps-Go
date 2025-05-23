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
 * Complete the 'hourglassSum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func hourglassSum(arr [][]int32) int32 {
	var maxSum int32 = -63 // The lowest possible sum (-9 * 7 elements)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// Calculate sum of the current hourglass
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + // Top row
				arr[i+1][j+1] + // Middle element
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2] // Bottom row

			// Update maxSum if we found a larger sum
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func main() {
	//I added
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	reader := bufio.NewReaderSize(file, 16*1024*1024)
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
