package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func lockit() (*os.File, error) {
	// Open the lock file
	lockFile, err := os.OpenFile("/data/liyang/lockfile", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	// Try to acquire a non-blocking exclusive lock
	err = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		// If locking fails, wait for lock to release
		fmt.Println("waiting for lock to release...")
		err = syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX)
		if err != nil {
			return nil, err
		}
	}

	return lockFile, nil
}

func main() {
	// Call lockit function to acquire the lock
	lockFile, err := lockit()
	if err != nil {
		fmt.Println("Error locking file:", err)
		return
	}
	defer lockFile.Close() // Ensure we release the lock when the program exits

	// Loop indefinitely and print "liyang" every second
	for {
		fmt.Println("liyang")
		time.Sleep(1 * time.Second)
	}
}

