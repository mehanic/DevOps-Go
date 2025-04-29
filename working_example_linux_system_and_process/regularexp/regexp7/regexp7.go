package main

import (
//	"errors"
	"fmt"
	"regexp"
)

// InvalidPhoneNumber is a custom error type.
type InvalidPhoneNumber struct {
	phoneNumber string
}

func (e *InvalidPhoneNumber) Error() string {
	return fmt.Sprintf("Invalid phone number: %s", e.phoneNumber)
}

var phonePattern = regexp.MustCompile(`(?x)
    ^            # match beginning of string
    \D*          # swallow anything that isn't numeric
    1?           # swallow leading 1, if present
    \D*          # swallow anything that isn't numeric
    (\d{3})      # capture 3-digit area code
    \D*          # swallow anything that isn't numeric
    (\d{3})      # capture 3-digit trunk
    \D*          # swallow anything that isn't numeric
    (\d{4})      # capture 4-digit number
    \D*          # swallow anything that isn't numeric
    (\d*)        # capture extension, if present
`)

// parsePhoneNumber parses a phone number string and returns the components.
func parsePhoneNumber(phoneNumber string) ([]string, error) {
	matches := phonePattern.FindStringSubmatch(phoneNumber)
	if matches != nil {
		return matches[1:], nil
	}
	return nil, &InvalidPhoneNumber{phoneNumber}
}

func main() {
	phoneNumbers := []string{
		"123-456-7890",
		"1-800-555-1212",
		"555-1234",
	}

	for _, number := range phoneNumbers {
		parsed, err := parsePhoneNumber(number)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Parsed phone number: %v\n", parsed)
		}
	}
}

