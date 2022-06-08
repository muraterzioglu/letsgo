package first_module

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func HelloAge(name string, age string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf("Hi, %v. You are %s years old today!", name, age)
	// nil means no error in go
	return message, nil
}

// In Go, the := operator is a shortcut for declaring and initializing a variable in one line
