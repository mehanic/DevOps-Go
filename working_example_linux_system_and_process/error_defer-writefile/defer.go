package main

import (
	"fmt"
	"os"

	"bufio"

	//"imooc-demo/functional/fib"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			// Uncomment panic to see
			// how it works with defer
			// panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}


// 99
// 98
// 97
// 96
// 95
// 94
// 93
// 92
// 91
// 90
// 89
// 88
// 87
// 86
// 85
// 84
// 83
// 82
// 81
// 80
// 79
// 78
// 77
// 76
// 75
// 74
// 73
// 72
// 71
// 70
// 69
// 68
// 67
// 66
// 65
// 64
// 63
// 62
// 61
// 60
// 59
// 58
// 57
// 56
// 55
// 54
// 53
// 52
// 51
// 50
// 49
// 48
// 47
// 46
// 45
// 44
// 43
// 42
// 41
// 40
// 39
// 38
// 37
// 36
// 35
// 34
// 33
// 32
// 31
// 30
// 29
// 28
// 27
// 26
// 25
// 24
// 23
// 22
// 21
// 20
// 19
// 18
// 17
// 16
// 15
// 14
// 13
// 12
// 11
// 10
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
