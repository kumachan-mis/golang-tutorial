package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hellos(names []string) (map[string]string, error) {
	messageMap := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messageMap[name] = message
	}
	return messageMap, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name error")
	}
	format := randomFormat()
	return fmt.Sprintf(format, name), nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	formatIndex := rand.Intn(len(formats))

	return formats[formatIndex]
}
