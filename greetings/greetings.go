package greetings

import (
	"errors"
	"fmt"
)

func HelloThere(name string) (string, error) {

	if "" == name {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
