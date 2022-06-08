package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

func WithRandom(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
