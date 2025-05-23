package main

import (
    "os/exec"
    "log"
)

func main() {
    cmd := exec.Command("jmeter", "-n", "-t", "test_plan.jmx", "-l", "results.jtl")
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Failed to run JMeter: %s", err)
    }
    log.Println("JMeter test completed successfully")
}
