package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSig sets up a listener for
// SIGINT interrupts
func CatchSig(ch chan os.Signal, done chan bool) {
	// block on waiting for a signal
	sig := <-ch
	// print it when it's received
	fmt.Println("\nsig received:", sig)

	// we can set up handlers for all types of
	// sigs here
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}

	// terminate
	done <- true
}

func main() {
	// initialize our channels
	signals := make(chan os.Signal)
	done := make(chan bool)

	// hook them up to the signals lib
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// if a signal is caught by this go routine
	// it will write to done
	go CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	// the program blocks until someone writes to done
	<-done
	fmt.Println("Done!")

}

//
//$./signals
//Press ctrl-c to terminate...
//^C
//sig received: interrupt
//handling a SIGINT now!
//Done!

//
//$./signals
//Press ctrl-c to terminate...
//# in a separate terminal
//$ ps -ef | grep signals
//501 30777 26360 0 5:00PM ttys000 0:00.00 ./signals
//$ kill -SIGTERM 30777
//# in the original terminal
//sig received: terminated
//handling a SIGTERM in an entirely different way!
//Done!
