package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// hello returns a greeting for the named person
func Hello(name string) (string, error) {

	// if no name was given, return an erro with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())

	return message, nil
}

// hellos returns a map that asssociates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error){

	// a map to associeate names with messages.
	messages := make(map[string]string)
	
	// loop through the received slice of names, calling
	// the hello funciton to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

	// in the map, associate the retrieved message with
	// the name

		messages[name] = message
	}

	return messages, nil
}

// init sets initial values for variable sused int the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomformat returns one of a set of gretting messages. the returned
// message is selecte a t random.
func randomFormat() string {
	// a slice of message formats
	formats := []string {
		"Hi, %v. Welcome",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	// return one of the message formats selected at random.
	return formats[rand.Intn(len(formats))]
}
