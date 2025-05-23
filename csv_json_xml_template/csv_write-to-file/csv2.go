package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

func main() {
    // Uncomment the following lines to read from a CSV file
    /*
        reqFile := "C:\\Users\\Automation\\Desktop\\hi\\new_info.csv"
        fo, err := os.Open(reqFile)
        if err != nil {
            fmt.Println("Error opening file:", err)
            return
        }
        defer fo.Close()

        csvReader := csv.NewReader(fo)
        csvReader.Comma = '|'
        for {
            eachRow, err := csvReader.Read()
            if err != nil {
                break
            }
            fmt.Println(eachRow)
        }
    */

    reqFile := "demo.csv"
    fo, err := os.Create(reqFile)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer fo.Close()

    csvWriter := csv.NewWriter(fo)
    csvWriter.Comma = ','

    // Uncomment the following lines to write individual rows
    /*
        csvWriter.Write([]string{"S_No", "Name", "Age"})
        csvWriter.Write([]string{"1", "John", "23"})
        csvWriter.Write([]string{"2", "Cliton", "24"})
    */

    myData := [][]string{
        {"S_No", "Name", "Age"},
        {"1", "John", "23"},
        {"2", "Cliton", "24"},
    }

    err = csvWriter.WriteAll(myData)
    if err != nil {
        fmt.Println("Error writing to CSV:", err)
    }

    csvWriter.Flush()
}

