package main

import (
    "fmt"
    "io"
    "strings"
)

// Reader interface already exists in io package
// We do not need to redefine it

// Writer interface
type Writer interface {
    Write(r []byte) (n int, err error)
}

// NotHowReaderIsDefined interface (incorrect example)
type NotHowReaderIsDefined interface {
    Read() (n []byte, err error)
}

// CountLetters counts the occurrences of each letter in the provided reader
func CountLetters(r io.Reader) (map[string]int, error) {
    buf := make([]byte, 2048)
    out := map[string]int{}
    for {
        n, err := r.Read(buf)
        for _, b := range buf[:n] {
            if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
                out[string(b)]++
            }
        }
        if err == io.EOF {
            return out, nil
        }
        if err != nil {
            return nil, err
        }
    }
}

func main() {
    text := "Hello, world! This is an example text to count letters."
    reader := strings.NewReader(text)

    letterCounts, err := CountLetters(reader)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Letter counts:", letterCounts)
}

//Letter counts: map[H:1 T:1 a:2 c:1 d:1 e:6 h:1 i:2 l:5 m:1 n:2 o:4 p:1 r:2 s:3 t:6 u:1 w:1 x:2]
