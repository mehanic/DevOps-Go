/* ASN1 DaytimeServer */

package main

import (
	"encoding/asn1"
	"log"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		log.Fatalf("Ошибка ResolveTCPAddr %v", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("Ошибка ListenTCP %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now()
		// Ignore return network errors.
		mdata, _ := asn1.Marshal(daytime)
		log.Print("Data ", mdata)
		conn.Write(mdata)
		conn.Close() // we're finished
		log.Print("Data send to ", conn.RemoteAddr())
	}
}

// go run dataTimeASNServer.go 

// 2025/01/15 21:48:59 Data [23 17 50 53 48 49 49 53 50 49 52 56 53 57 43 48 49 48 48]
// 2025/01/15 21:48:59 Data send to 127.0.0.1:56942
// 2025/01/15 21:51:00 Data [23 17 50 53 48 49 49 53 50 49 53 49 48 48 43 48 49 48 48]
// 2025/01/15 21:51:00 Data send to 127.0.0.1:58786
