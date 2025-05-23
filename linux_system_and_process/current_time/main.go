package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
	// Run 'go get github.com/beevik/ntp' to install the package	"os"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Errorf("error: %v", err)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time %s", time)
}

//Current time 2025-01-15 18:53:57.923564874 +0100 CET m=+0.524201217