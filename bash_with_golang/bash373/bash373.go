package main

import (
	"fmt"
	"os"
	"syscall"
//	"time"
)

func main() {
	// Check if /tmp/file1 exists
	fmt.Print("Does file /tmp/file1 exist? ")
	if _, err := os.Stat("/tmp/file1"); os.IsNotExist(err) {
		fmt.Println(1) // Equivalent to false in Bash
	} else {
		fmt.Println(0) // Equivalent to true in Bash
	}

	// Create /tmp/file1 and /tmp/file2
	os.Create("/tmp/file1")
	os.Create("/tmp/file2")

	// Check if /tmp/file1 exists after creation
	fmt.Print("Does file /tmp/file1 exist now? ")
	if _, err := os.Stat("/tmp/file1"); os.IsNotExist(err) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}

	// Check if /tmp is a directory
	fmt.Print("Is /tmp a directory? ")
	if info, err := os.Stat("/tmp"); err == nil && info.IsDir() {
		fmt.Println(0) // true
	} else {
		fmt.Println(1) // false
	}

	// Check if sticky bit is set on /tmp
	fmt.Print("Is sticky bit set on /tmp? ")
	var stat syscall.Stat_t
	if err := syscall.Stat("/tmp", &stat); err == nil && stat.Mode&syscall.S_ISVTX != 0 {
		fmt.Println(0) // true
	} else {
		fmt.Println(1) // false
	}

	// Check if /tmp has execute permission
	fmt.Print("Does /tmp have execute permission? ")
	if info, err := os.Stat("/tmp"); err == nil && info.Mode()&0111 != 0 {
		fmt.Println(0) // true
	} else {
		fmt.Println(1) // false
	}

	// Create /tmp/file2
	os.Create("/tmp/file2")

	// Check if /tmp/file1 is newer than /tmp/file2
	fmt.Print("Is /tmp/file1 modified more recently than /tmp/file2? ")
	file1Info, err1 := os.Stat("/tmp/file1")
	file2Info, err2 := os.Stat("/tmp/file2")

	if err1 == nil && err2 == nil && file1Info.ModTime().After(file2Info.ModTime()) {
		fmt.Println(0) // true
	} else {
		fmt.Println(1) // false
	}
}

