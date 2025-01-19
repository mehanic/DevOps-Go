package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

func main() {
	// version 1 uuid
	v1, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("cannot generate v1 uuid")
	}
	fmt.Printf("v1 uuid: %v\n", v1)

	// version 2 uuid
	v2, err := uuid.NewDCEGroup()
	if err != nil {
		log.Fatal("cannot generate v2 uuid")
	}
	fmt.Printf("v2 uuid: %v\n", v2)

	// version 3 uuid
	v3 := uuid.NewMD5(uuid.NameSpaceURL, []byte("https://example.com"))
	fmt.Printf("v3 uuid: %v\n", v3)

	// version 4 uuid
	v4, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("cannot generate v4 uuid")
	}
	fmt.Printf("v4 uuid: %v\n", v4)

	// version 5 uuid
	v5 := uuid.NewSHA1(uuid.NameSpaceURL, []byte("https://example.com"))
	fmt.Printf("v5 uuid: %v\n", v5)

	
}

// v1 uuid: a895f64b-d698-11ef-ad8d-98fa9bbd8c21
// v2 uuid: 000003e8-d698-21ef-ad01-98fa9bbd8c21
// v3 uuid: 68794df6-5e20-385f-ab08-bb73f8a433cb
// v4 uuid: 03824062-6d60-47fc-b393-e61b2bdeb01d
// v5 uuid: 4fd35a71-71ef-5a55-a9d9-aa75c889a6d0
