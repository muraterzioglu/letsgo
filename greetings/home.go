package greetings

import (
	"errors"
	"fmt"
)

// JustName returns a greeting for the named person.
func JustName(name string) string {
	// Return a greeting that embeds the name in a message.
	// In Go, the := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

/*
	if function starts with upper case its exported
	from that module.
*/
func WithAge(name string, age string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf("Hi, %v. You are %s years old today!", name, age)
	// nil means no error in go
	return message, nil
}
