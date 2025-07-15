package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "regexp"
    "sort"
)


/*
 * Complete the 'mostUsedPrefixes' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY numbers as parameter.
 */

func mostUsedPrefixes(numbers []string) []string {
re := regexp.MustCompile(`^(\d{3})-(\d{3})-(\d{4})$`)
    counts := make(map[string]int)

    for _, number := range numbers {
        if match := re.FindStringSubmatch(number); match != nil {
            areaCode := match[1]
            counts[areaCode]++
        }
    }

    type pair struct {
        code  string
        count int
    }

    var list []pair
    for k, v := range counts {
        list = append(list, pair{k, v})
    }

    sort.Slice(list, func(i, j int) bool {
        if list[i].count == list[j].count {
            return list[i].code < list[j].code
        }
        return list[i].count > list[j].count
    })

    top := []string{}
    for i := 0; i < len(list) && i < 3; i++ {
        top = append(top, list[i].code)
    }
    return top

}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    numbersCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var numbers []string

    for i := 0; i < int(numbersCount); i++ {
        numbersItem := readLine(reader)
        numbers = append(numbers, numbersItem)
    }

    result := mostUsedPrefixes(numbers)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
