
package main

import (
    "fmt"
    "time"
)

func main() {

    // Example #1

    requests := make(chan int, 15)

    for i := 0; i < 5; i++ {
        requests <- 1
    }

    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <- limiter
        fmt.Println("request", req, time.Now())
    }

    // Example #2

    plosiveRequests := make(chan int, 5)

    for i := 0; i < 5; i++ {
        plosiveRequests <- i
    }

    close(plosiveRequests)

    plosiveLimiter := make(chan time.Time, 3)

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            plosiveLimiter <- t
        }
    }()

    for req := range plosiveRequests {
        <-plosiveLimiter
        fmt.Println("plosive request", req, time.Now())
    }
}


// request 1 2025-02-14 15:15:49.112510108 +0100 CET m=+0.200437855
// request 1 2025-02-14 15:15:49.31294229 +0100 CET m=+0.400870037
// request 1 2025-02-14 15:15:49.512279873 +0100 CET m=+0.600207619
// request 1 2025-02-14 15:15:49.712665589 +0100 CET m=+0.800593335
// request 1 2025-02-14 15:15:49.913044125 +0100 CET m=+1.000971871
// plosive request 0 2025-02-14 15:15:50.113409364 +0100 CET m=+1.201337113
// plosive request 1 2025-02-14 15:15:50.313736985 +0100 CET m=+1.401664733
// plosive request 2 2025-02-14 15:15:50.514024473 +0100 CET m=+1.601952220
// plosive request 3 2025-02-14 15:15:50.713292644 +0100 CET m=+1.801220391
// plosive request 4 2025-02-14 15:15:50.913577229 +0100 CET m=+2.001504977
