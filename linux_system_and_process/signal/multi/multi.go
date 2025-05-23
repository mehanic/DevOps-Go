package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
)

// Redefine Unix only signals, so this compiles under Windows
const (
	SIGUSR1 = syscall.Signal(0xa)
	SIGUSR2 = syscall.Signal(0xc)
)

var cfgPath string

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatalln("user:", err)
	}
	cfgPath = filepath.Join(u.HomeDir, ".multi")
}

func main() {
	c := make(chan os.Signal, 1)
	d := time.Second
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGQUIT, SIGUSR1, SIGUSR2, syscall.SIGALRM)
	// initial load
	if err := handleSignal(syscall.SIGHUP, &d); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	for {
		select {
		case s := <-c:
			if err := handleSignal(s, &d); err != nil {
				log.Printf("Error handling %s: %s", s, err)
				continue
			}
		default:
			time.Sleep(d)
			log.Println("After", d, "Executing action!")
		}
	}
}

func handleSignal(s os.Signal, d *time.Duration) error {
	switch s {
	case syscall.SIGHUP:
		return loadSettings(d)
	case syscall.SIGALRM:
		return saveSettings(d)
	case syscall.SIGINT:
		if err := saveSettings(d); err != nil {
			log.Println("Cannot save:", err)
			os.Exit(1)
		}
		fallthrough
	case syscall.SIGQUIT:
		os.Exit(0)
	case SIGUSR1:
		changeSettings(d, (*d)*2)
		return nil
	case SIGUSR2:
		changeSettings(d, (*d)/2)
		return nil
	}
	return nil
}

func changeSettings(d *time.Duration, v time.Duration) {
	*d = v
	log.Println("Changed", v)
}

func loadSettings(d *time.Duration) error {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return err
	}
	var v time.Duration
	if v, err = time.ParseDuration(string(b)); err != nil {
		return err
	}
	*d = v
	log.Println("Loaded", v)
	return nil
}

func saveSettings(d *time.Duration) error {
	f, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = fmt.Fprint(f, d); err != nil {
		return err
	}
	log.Println("Saved", *d)
	return nil
}


// └> $ go run multi.go 
// 2024/02/27 01:55:44 Loaded 1s
// 2024/02/27 01:55:45 After 1s Executing action!
// 2024/02/27 01:55:46 After 1s Executing action!
// 2024/02/27 01:55:47 After 1s Executing action!
// 2024/02/27 01:55:48 After 1s Executing action!
// 2024/02/27 01:55:49 After 1s Executing action!
// 2024/02/27 01:55:50 After 1s Executing action!
// 2024/02/27 01:55:51 After 1s Executing action!
// 2024/02/27 01:55:52 After 1s Executing action!
// 2024/02/27 01:55:53 After 1s Executing action!
// 2024/02/27 01:55:54 After 1s Executing action!
// 2024/02/27 01:55:55 After 1s Executing action!
// ^C2024/02/27 01:55:56 After 1s Executing action!
// 2024/02/27 01:55:56 Saved 1s
