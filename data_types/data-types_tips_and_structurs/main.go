package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type (
	ID        = uint64
	UUIDAlias = [32]byte
	UserIDs   = []ID
	bytes     = []byte
)

func aliasTypes() {
	var myID ID = 1
	myIDUint64 := myID
	fmt.Println("Alias Type Example - myID as uint64:", myIDUint64)
}

type IDWrapper uint64

func (id IDWrapper) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

type UUID [32]byte

func (id UUID) String() string {
	sb := strings.Builder{}
	for _, b := range id {
		sb.WriteByte(b)
	}
	return sb.String()
}

func uuidExample() {
	// Demonstrate JSON unmarshalling
	jsonData := []uint8{72, 101, 108, 108, 111}
	var result []byte
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	} else {
		fmt.Println("Unmarshalled data as bytes:", result)
	}
}

type (
	Interface interface {
		Method()
	}

	Structure struct {
		ID    uint64 `tag1:"tagValue1|tagValue2|tagValue3" tag2:"tagValue" tag3:"tagValue" gorm:"<-create;<-update;<-delete"`
		Value string `binding:"email"` // go-validator
	}

	Wrapper     uint64
	WrapperFunc func()
)

func main() {
	// Alias type example
	aliasTypes()

	// Wrapper type example
	var wrappedID IDWrapper = 12345
	fmt.Println("Wrapper Type Example - IDWrapper:", wrappedID.String())

	// UUID example
	var uuid UUID
	copy(uuid[:], "GoLangUUIDExampleDataHere...")
	fmt.Println("UUID Example - UUID as String:", uuid.String())

	// JSON unmarshalling example
	uuidExample()

	// Structure and Tags example
	user := Structure{
		ID:    1001,
		Value: "user@example.com",
	}
	fmt.Printf("Structure Example - User ID: %d, Value: %s\n", user.ID, user.Value)
}

// Alias Type Example - myID as uint64: 1
// Wrapper Type Example - IDWrapper: 12345
// UUID Example - UUID as String: GoLangUUIDExampleDataHere...
// Error unmarshalling JSON: invalid character 'H' looking for beginning of value
// Structure Example - User ID: 1001, Value: user@example.com
