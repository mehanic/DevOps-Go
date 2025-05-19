package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

// scan function checks the specified ports on the given host
func scan(host string) {
    ports := []int{80, 110, 443, 8080}
    for _, port := range ports {
        address := fmt.Sprintf("%s:%d", host, port)
        conn, err := net.DialTimeout("tcp", address, 3*time.Second) // Timeout set to 3 seconds
        if err != nil {
            fmt.Printf("Port %d: Closed or filtered\n", port)
        } else {
            fmt.Printf("Port %d: Open\n", port)
            conn.Close() // Close the connection if it was successful
        }
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("[+] Usage: go run main.go 192.168.1.1")
        os.Exit(1)
    }
    host := os.Args[1]
    scan(host)
}

