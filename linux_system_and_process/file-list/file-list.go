package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := os.ReadDir("testFolder")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	main1()
	main2()
    main3()
	main4()
	main5()
	main6()	
}

//---

func main1() {
	err := filepath.Walk("/tmp/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

//-

func main2() {
	files, err := os.ReadDir("/tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}

//--

func main3() {
	dir, err := os.Open("testFolder")
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

//--

func main4() {
	err := filepath.Walk("testFolder",
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}
}

//--

func main5() {
	files, err := filepath.Glob("testFolder/*.go")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}

//--

func main6() {
	f, err := os.Open("/tmp")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		fmt.Println(v.Name(), v.IsDir())
	}
}

//
