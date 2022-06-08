package greetings

import (
	"math/rand"
	"time"
)

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the JustName function to get a message for each name.

	// _ => key of the loop
	// name => item name of the map
	for _, name := range names {
		message, err := WithRandom(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}
