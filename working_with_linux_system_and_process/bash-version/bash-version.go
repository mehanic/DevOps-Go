package main

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    cmd := exec.Command("bash", "--version")

    var out bytes.Buffer
    var errOut bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &errOut

    err := cmd.Run()
    if err != nil {
        fmt.Println("Command was failed and error is:", errOut.String())
        return
    }

    output := out.String()
    for _, eachLine := range strings.Split(output, "\n") {
        if strings.Contains(eachLine, "version") && strings.Contains(eachLine, "release") {
            parts := strings.Fields(eachLine)
            if len(parts) > 3 {
                version := strings.Split(parts[3], "(")[0]
                fmt.Println(version)
            }
        }
    }
}
