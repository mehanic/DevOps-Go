package error_handling

import (
	"fmt"
	"time"
)

type AgeLimitException struct {
	birthyear int
	message   string
}

type WorkingHourException struct {
	hour    int
	message string
}

func (e *AgeLimitException) Error() string {
	return fmt.Sprintf("%d --- %s", e.birthyear, e.message)
}

func (e *WorkingHourException) Error() string {
	return fmt.Sprintf("%d --- %s", e.hour, e.message)
}

func EnterCheck(birthyear int, hour int) (string, error) {

	year, _, _ := time.Now().Date()

	age := year - birthyear

	if hour > 11 && hour < 24 {
		if age >= 18 {
			return "You're allowed to enter. Have Fun!", nil
		} else {
			return "", &AgeLimitException{age, "You should be at least 18 years old to enter!"}
		}
	} else {
		return "", &WorkingHourException{hour, "Sorry, we work between 11am and 12 pm"}
	}

}
