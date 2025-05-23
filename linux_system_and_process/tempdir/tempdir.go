package main

import (
	"fmt"
	"io"
	"os"
)

func workWithTemporaryFile() {
	tf, err := os.CreateTemp("", "example")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(tf.Name())

	if _, err := tf.Write([]byte("Order Id 93827, Account Id 286438273")); err != nil {
		fmt.Println(err)
		return
	}

	if _, err := tf.Seek(0, 0); err != nil {
		fmt.Println(err)
		return
	}

	data, err := io.ReadAll(tf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
	fmt.Println(os.TempDir())
}
func main() {
	workWithTemporaryFile()
}
